package ptestingclient_test

import (
	"context"
	"github.com/basemind-ai/monorepo/e2e/factories"
	"github.com/basemind-ai/monorepo/gen/go/ptesting/v1"
	"github.com/basemind-ai/monorepo/services/dashboard-backend/internal/dto"
	"github.com/basemind-ai/monorepo/services/dashboard-backend/internal/ptestingclient"
	"github.com/basemind-ai/monorepo/shared/go/db"
	"github.com/basemind-ai/monorepo/shared/go/testutils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

func TestMain(m *testing.M) {
	cleanup := testutils.CreateNamespaceTestDBModule("grpc-client-test")
	defer cleanup()
	m.Run()
}

func CreateClientAndService(
	t *testing.T,
) (*ptestingclient.Client, *testutils.MockPromptTestingService) {
	t.Helper()
	mockService := &testutils.MockPromptTestingService{T: t}
	listener := testutils.CreateTestGRPCServer[ptesting.PromptTestingServiceServer](
		t,
		ptesting.RegisterPromptTestingServiceServer,
		mockService,
	)
	client := ptestingclient.New(
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			},
		),
	)

	return client, mockService
}

func TestPromptTestingGRPCClient(t *testing.T) {
	testutils.SetTestEnv(t)

	project, _ := factories.CreateProject(context.TODO())
	application, _ := factories.CreateApplication(context.TODO(), project.ID)
	promptConfig, _ := factories.CreatePromptConfig(context.TODO(), application.ID)
	promptRequestRecord, _ := factories.CreatePromptRequestRecord(context.TODO(), promptConfig.ID)

	applicationID := db.UUIDToString(&application.ID)
	promptConfigID := db.UUIDToString(&promptConfig.ID)
	promptRequestRecordID := db.UUIDToString(&promptRequestRecord.ID)
	data := dto.PromptConfigTestDTO{
		Name:                   "TEST",
		ModelParameters:        promptConfig.ModelParameters,
		ModelType:              db.ModelTypeGpt432k,
		ModelVendor:            db.ModelVendorOPENAI,
		ProviderPromptMessages: promptConfig.ProviderPromptMessages,
		TemplateVariables:      map[string]string{},
		PromptConfigID:         &promptConfigID,
	}

	t.Run("New", func(t *testing.T) {
		t.Run("dials and returns client", func(t *testing.T) {
			client := ptestingclient.New(
				"localhost:50051",
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			assert.NotNil(t, client)
		})
	})

	t.Run("Init", func(t *testing.T) {
		t.Run("panics if the env is not set", func(t *testing.T) {
			assert.Panics(t,
				func() {
					ptestingclient.Init(
						context.Background(),
						grpc.WithTransportCredentials(insecure.NewCredentials()),
					)
				},
			)
		})
		t.Run("does not panic if the env is set", func(t *testing.T) {
			t.Setenv("API_GATEWAY_ADDRESS", "localhost:50051")
			assert.NotPanics(t,
				func() {
					ptestingclient.Init(
						context.Background(),
						grpc.WithTransportCredentials(insecure.NewCredentials()),
					)
				},
			)
		})
	})

	t.Run("GetClient", func(t *testing.T) {
		t.Run("does not panic if init is called", func(t *testing.T) {
			t.Setenv("API_GATEWAY_ADDRESS", "localhost:50051")
			ptestingclient.Init(
				context.Background(),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			assert.NotPanics(t, func() {
				ptestingclient.GetClient()
			})
		})
		t.Run("panics if init is not called", func(t *testing.T) {
			ptestingclient.SetClient(nil)
			assert.Panics(t, func() {
				ptestingclient.GetClient()
			})
		})
	})

	t.Run("handles a stream response", func(t *testing.T) {
		client, mockService := CreateClientAndService(t)
		responseChannel := make(chan *ptesting.PromptTestingStreamingPromptResponse)
		errorChannel := make(chan error)

		finishReason := "done"

		mockService.Stream = []*ptesting.PromptTestingStreamingPromptResponse{
			{Content: "1"},
			{Content: "2"},
			{
				Content:               "3",
				FinishReason:          &finishReason,
				PromptRequestRecordId: &promptRequestRecordID,
			},
		}

		go client.StreamPromptTest(
			context.TODO(),
			applicationID,
			&data,
			responseChannel,
			errorChannel,
		)

		chunks := make([]*ptesting.PromptTestingStreamingPromptResponse, 0)

		for chunk := range responseChannel {
			chunks = append(chunks, chunk)
		}

		assert.Len(t, chunks, 3)
		assert.Equal(t, "1", chunks[0].Content)
		assert.Nil(t, chunks[0].FinishReason)
		assert.Nil(t, chunks[0].PromptRequestRecordId)
		assert.Equal(t, "2", chunks[1].Content)
		assert.Nil(t, chunks[1].FinishReason)
		assert.Nil(t, chunks[1].PromptRequestRecordId)
		assert.Equal(t, "3", chunks[2].Content)
		assert.Equal(t, finishReason, *chunks[2].FinishReason)
		assert.Equal(t, promptRequestRecordID, *chunks[2].PromptRequestRecordId)
	})
	t.Run("sends error using the error channel on stream error", func(t *testing.T) {
		client, mockService := CreateClientAndService(t)
		responseChannel := make(chan *ptesting.PromptTestingStreamingPromptResponse)
		errorChannel := make(chan error)

		mockService.Error = assert.AnError

		go client.StreamPromptTest(
			context.TODO(),
			applicationID,
			&data,
			responseChannel,
			errorChannel,
		)

	loop:
		for {
			select {
			case err := <-errorChannel:
				assert.Error(t, err)
				break loop

			case <-responseChannel:
				t.Fatal("should not have received data")
			}
		}
	})
}
