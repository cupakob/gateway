package service

import (
	"context"
	"github.com/basemind-ai/monorepo/gen/go/gateway/v1"
	"github.com/basemind-ai/monorepo/go-services/api-gateway/connectors"
	"github.com/basemind-ai/monorepo/go-services/api-gateway/constants"
	"github.com/basemind-ai/monorepo/go-shared/db"
	"github.com/basemind-ai/monorepo/go-shared/rediscache"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
	gateway.UnimplementedAPIGatewayServiceServer
}

func New() gateway.APIGatewayServiceServer {
	return Server{}
}

func RetrieveApplicationHandler(ctx context.Context, applicationId string) func() (*db.Application, error) {
	return func() (*db.Application, error) {
		appId := pgtype.UUID{}
		if idErr := appId.Scan(applicationId); idErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid application id: %v", idErr)
		}

		application, applicationQueryErr := db.GetQueries().FindApplicationById(ctx, appId)

		if applicationQueryErr != nil {
			return nil, status.Errorf(codes.NotFound, "application with the given ID does not exist: %v", applicationQueryErr)
		}
		return &application, nil
	}
}

func (Server) RequestPromptConfig(ctx context.Context, _ *gateway.PromptConfigRequest) (*gateway.PromptConfigResponse, error) {
	applicationId, ok := ctx.Value(constants.ApplicationIDContextKey).(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "application ID not found in context")
	}

	application, retrievalErr := rediscache.With[db.Application](
		ctx, applicationId, &db.Application{}, time.Minute*30, RetrieveApplicationHandler(ctx, applicationId),
	)
	if retrievalErr != nil {
		return nil, retrievalErr
	}

	return &gateway.PromptConfigResponse{
		ExpectedPromptVariables: application.ExpectedTemplateVariables,
	}, nil
}

func (Server) RequestPrompt(ctx context.Context, request *gateway.PromptRequest) (*gateway.PromptResponse, error) {
	applicationId, ok := ctx.Value(constants.ApplicationIDContextKey).(string)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "application ID not found in context")
	}

	application, retrievalErr := rediscache.With[db.Application](
		ctx, applicationId, &db.Application{}, time.Minute*30, RetrieveApplicationHandler(ctx, applicationId),
	)
	if retrievalErr != nil {
		return nil, retrievalErr
	}

	client := connectors.GetOpenAIConnectorClient()
	responseContent, requestErr := client.RequestPrompt(ctx, applicationId, *application, request.TemplateVariables)
	if requestErr != nil {
		return nil, retrievalErr
	}

	return &gateway.PromptResponse{
		Content:      responseContent,
		PromptTokens: 0,
	}, nil
}
func (Server) RequestStreamingPrompt(request *gateway.PromptRequest, streamServer gateway.APIGatewayService_RequestStreamingPromptServer) error {
	applicationId, ok := streamServer.Context().Value(constants.ApplicationIDContextKey).(string)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "application ID not found in context")
	}

	application, retrievalErr := rediscache.With[db.Application](
		streamServer.Context(), applicationId, &db.Application{}, time.Minute*30, RetrieveApplicationHandler(streamServer.Context(), applicationId),
	)
	if retrievalErr != nil {
		return retrievalErr
	}

	contentChannel := make(chan string)
	errorChannel := make(chan error, 1)

	go connectors.GetOpenAIConnectorClient().RequestStream(streamServer.Context(), applicationId, *application, request.TemplateVariables, contentChannel, errorChannel)

	for {
		select {
		case content := <-contentChannel:
			if sendErr := streamServer.SendMsg(&gateway.StreamingPromptResponse{
				Content: content,
			}); sendErr != nil {
				return sendErr
			}
		case err := <-errorChannel:
			return err
		}
	}
}
