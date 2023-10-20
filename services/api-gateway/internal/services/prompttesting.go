package services

import (
	"fmt"
	"github.com/basemind-ai/monorepo/gen/go/ptesting/v1"
	"github.com/basemind-ai/monorepo/services/api-gateway/internal/connectors"
	"github.com/basemind-ai/monorepo/services/api-gateway/internal/dto"
	"github.com/basemind-ai/monorepo/shared/go/datatypes"
	"github.com/basemind-ai/monorepo/shared/go/db"
)

type PromptTestingServer struct {
	ptesting.UnimplementedPromptTestingServiceServer
}

func (PromptTestingServer) TestPrompt(
	request *ptesting.PromptTestRequest,
	streamServer ptesting.PromptTestingService_TestPromptServer,
) error {
	channel := make(chan dto.PromptResultDTO)

	applicationID, applicationIDParseErr := db.StringToUUID(request.ApplicationId)
	if applicationIDParseErr != nil {
		return fmt.Errorf("failed to parse application ID: %w", applicationIDParseErr)
	}

	promptConfigID, promptConfigIDParseErr := db.StringToUUID(request.PromptConfigId)
	if promptConfigIDParseErr != nil {
		return fmt.Errorf("failed to parse prompt config ID: %w", promptConfigIDParseErr)
	}

	requestConfigurationDTO := &dto.RequestConfigurationDTO{
		ApplicationIDString: request.ApplicationId,
		ApplicationID:       *applicationID,
		PromptConfigID:      *promptConfigID,
		PromptConfigData: datatypes.PromptConfigDTO{
			ID:                        request.PromptConfigId,
			ModelParameters:           request.ModelParameters,
			ModelType:                 db.ModelType(request.ModelType),
			ModelVendor:               db.ModelVendor(request.ModelVendor),
			ProviderPromptMessages:    request.ProviderPromptMessages,
			ExpectedTemplateVariables: request.ExpectedTemplateVariables,
		},
	}

	go connectors.GetProviderConnector(db.ModelVendor(request.ModelVendor)).
		RequestStream(
			streamServer.Context(),
			requestConfigurationDTO,
			request.TemplateVariables,
			channel,
		)

	return streamFromChannel(channel, streamServer, createPromptTestingStreamMessage)
}

func createPromptTestingStreamMessage(
	result dto.PromptResultDTO,
) (*ptesting.PromptTestingStreamingPromptResponse, bool) {
	msg := &ptesting.PromptTestingStreamingPromptResponse{}
	if result.Error != nil {
		reason := "error"
		msg.FinishReason = &reason
	}
	if result.RequestRecord != nil {
		if msg.FinishReason == nil {
			reason := "done"
			msg.FinishReason = &reason
		}
		promptRequestRecordID := db.UUIDToString(&result.RequestRecord.ID)
		msg.PromptRequestRecordId = &promptRequestRecordID
	}

	if result.Content != nil {
		msg.Content = *result.Content
	}
	return msg, msg.FinishReason != nil
}
