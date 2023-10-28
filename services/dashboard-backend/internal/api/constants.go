package api

const (
	ApplicationAnalyticsEndpoint    = "/projects/{projectId}/applications/{applicationId}/analytics"
	ApplicationDetailEndpoint       = "/projects/{projectId}/applications/{applicationId}"
	ApplicationAPIKeyDetailEndpoint = "/projects/{projectId}/applications/{applicationId}/apikeys/{apiKeyId}" //nolint: gosec
	ApplicationAPIKeysListEndpoint  = "/projects/{projectId}/applications/{applicationId}/apikeys"            //nolint: gosec
	ApplicationsListEndpoint        = "/projects/{projectId}/applications"
	ProjectAnalyticsEndpoint        = "/projects/{projectId}/analytics"
	ProjectDetailEndpoint           = "/projects/{projectId}"
	ProjectOTPEndpoint              = "/projects/{projectId}/otp"
	ProjectUserDetailEndpoint       = "/projects/{projectId}/users/{userId}"
	ProjectUserListEndpoint         = "/projects/{projectId}/users"
	ProjectsListEndpoint            = "/projects"
	PromptConfigAnalyticsEndpoint   = "/projects/{projectId}/applications/{applicationId}/prompt-configs/{promptConfigId}/analytics"
	PromptConfigDetailEndpoint      = "/projects/{projectId}/applications/{applicationId}/prompt-configs/{promptConfigId}"
	PromptConfigListEndpoint        = "/projects/{projectId}/applications/{applicationId}/prompt-configs"
	PromptConfigSetDefaultEndpoint  = "/projects/{projectId}/applications/{applicationId}/prompt-configs/{promptConfigId}/set-default"
	PromptConfigTestingEndpoint     = "/projects/{projectId}/applications/{applicationId}/prompt-configs/test"
	UserAccountDetailEndpoint       = "/users"
)

const (
	invalidRequestBodyError = "invalid request body"
)
