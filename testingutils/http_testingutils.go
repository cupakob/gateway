package testingutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/basemind-ai/backend-services/lib/httpclient"
)

func CreateTestClient(t *testing.T, handler http.Handler) httpclient.Client {
	server := httptest.NewServer(handler)

	t.Cleanup(func() {
		server.Close()
	})

	return httpclient.New(server.URL, server.Client())
}
