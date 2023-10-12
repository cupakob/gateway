package api_test

import (
	"context"
	"github.com/basemind-ai/monorepo/services/dashboard-backend/internal/api"
	"github.com/basemind-ai/monorepo/services/dashboard-backend/internal/middleware"
	"github.com/basemind-ai/monorepo/shared/go/httpclient"
	httpTestUtils "github.com/basemind-ai/monorepo/shared/go/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/basemind-ai/monorepo/e2e/factories"
	"github.com/basemind-ai/monorepo/shared/go/db"
	"github.com/basemind-ai/monorepo/shared/go/router"
)

func TestMain(m *testing.M) {
	cleanup := httpTestUtils.CreateNamespaceTestDBModule("api-test")
	defer cleanup()
	m.Run()
}

func createProject(t *testing.T) string {
	t.Helper()
	project, _ := factories.CreateProject(context.TODO())
	return db.UUIDToString(&project.ID)
}

func createUserProject(
	t *testing.T,
	firebaseID string,
	projectID string,
	permission db.AccessPermissionType,
) {
	t.Helper()

	userAccount, err := db.
		GetQueries().
		RetrieveUserAccount(context.TODO(), firebaseID)

	assert.NoError(t, err)

	projectIDUUID, err := db.StringToUUID(projectID)
	assert.NoError(t, err)

	_, err = db.GetQueries().CreateUserProject(context.TODO(), db.CreateUserProjectParams{
		UserID:     userAccount.ID,
		ProjectID:  *projectIDUUID,
		Permission: permission,
	})
	assert.NoError(t, err)
}

func createApplication(t *testing.T, projectID string) string {
	t.Helper()
	uuidID, _ := db.StringToUUID(projectID)
	application, _ := factories.CreateApplication(context.TODO(), *uuidID)
	applicationID := db.UUIDToString(&application.ID)
	return applicationID
}

func createTestClient(t *testing.T, userAccount *db.UserAccount) httpclient.Client {
	t.Helper()
	r := router.New(router.Options{
		Environment:      "test",
		ServiceName:      "test",
		RegisterHandlers: api.RegisterHandlers,
		Middlewares: []func(next http.Handler) http.Handler{
			middleware.CreateMockFirebaseAuthMiddleware(userAccount),
		},
	})

	return httpTestUtils.CreateTestHTTPClient(t, r)
}
