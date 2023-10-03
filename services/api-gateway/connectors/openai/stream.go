package openai

import (
	"context"
	"errors"
	"fmt"
	openaiconnector "github.com/basemind-ai/monorepo/gen/go/openai/v1"
	"github.com/basemind-ai/monorepo/shared/go/datatypes"
	"github.com/basemind-ai/monorepo/shared/go/db"
	"github.com/basemind-ai/monorepo/shared/go/tokenutils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"io"
	"strings"
	"time"
)

func streamFromClient(
	channel chan<- datatypes.PromptResult,
	promptResult *datatypes.PromptResult,
	recordParams *db.CreatePromptRequestRecordParams,
	startTime time.Time,
	stream openaiconnector.OpenAIService_OpenAIStreamClient,
) string {
	var builder strings.Builder

	for {
		if msg, receiveErr := stream.Recv(); receiveErr != nil {
			recordParams.FinishTime = pgtype.Timestamptz{Time: time.Now(), Valid: true}

			if !errors.Is(receiveErr, io.EOF) {
				log.Debug().Err(receiveErr).Msg("received stream error")
				promptResult.Error = receiveErr
			}

			break
		} else {
			if recordParams.StreamResponseLatency.Int64 == 0 {
				duration := int64(time.Until(startTime))
				recordParams.StreamResponseLatency = pgtype.Int8{Int64: duration, Valid: true}
			}

			builder.WriteString(msg.Content)
			channel <- datatypes.PromptResult{Content: &msg.Content}
		}
	}

	return builder.String()
}

func (c *Client) RequestStream(
	ctx context.Context,
	applicationId string,
	requestConfiguration *datatypes.RequestConfiguration,
	templateVariables map[string]string,
	channel chan<- datatypes.PromptResult,
) {
	promptRequest, promptRequestErr := CreatePromptRequest(
		applicationId,
		requestConfiguration.PromptConfigData.ModelType,
		requestConfiguration.PromptConfigData.ModelParameters,
		requestConfiguration.PromptConfigData.ProviderPromptMessages,
		templateVariables,
	)
	if promptRequestErr != nil {
		log.Error().Err(promptRequestErr).Msg("failed to create prompt request")
		channel <- datatypes.PromptResult{Error: promptRequestErr}
		close(channel)

		return
	}

	startTime := time.Now()

	recordParams := &db.CreatePromptRequestRecordParams{
		PromptConfigID:   requestConfiguration.PromptConfigData.ID,
		IsStreamResponse: true,
		StartTime:        pgtype.Timestamptz{Time: startTime, Valid: true},
		RequestTokens: tokenutils.GetPromptTokenCount(
			GetRequestPromptString(promptRequest.Messages),
			requestConfiguration.PromptConfigData.ModelType,
		),
	}
	finalResult := &datatypes.PromptResult{}

	if stream, streamErr := c.client.OpenAIStream(ctx, promptRequest); streamErr == nil {
		promptContent := streamFromClient(
			channel,
			finalResult,
			recordParams,
			startTime,
			stream,
		)
		recordParams.ResponseTokens = tokenutils.GetPromptTokenCount(
			promptContent,
			requestConfiguration.PromptConfigData.ModelType,
		)
	} else {
		finalResult.Error = streamErr
	}

	if finalResult.Error != nil {
		recordParams.ErrorLog = pgtype.Text{String: finalResult.Error.Error(), Valid: true}
	}

	if promptRecord, createRequestRecordErr := db.GetQueries().CreatePromptRequestRecord(ctx, *recordParams); createRequestRecordErr != nil {
		log.Error().Err(createRequestRecordErr).Msg("failed to create prompt request record")
		if finalResult.Error == nil {
			finalResult.Error = createRequestRecordErr
		} else {
			finalResult.Error = fmt.Errorf("failed to save prompt record: %w...%w", finalResult.Error, createRequestRecordErr)
		}
	} else {
		finalResult.RequestRecord = &promptRecord
	}
	channel <- *finalResult
	close(channel)
}
