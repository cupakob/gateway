package dto

import (
	"encoding/json"
	"github.com/basemind-ai/monorepo/shared/go/db/models"
	"time"
)

// ApplicationDTO - DTO for serializing application data.
type ApplicationDTO struct { // skipcq: TCV-001
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ProjectDTO - DTO for serializing project data.
type ProjectDTO struct { // skipcq: TCV-001
	ID           string           `json:"id,omitempty"`
	Name         string           `json:"name"                   validate:"required"`
	Description  string           `json:"description,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
	Permission   string           `json:"permission,omitempty"`
	Applications []ApplicationDTO `json:"applications,omitempty"`
}

// PromptConfigCreateDTO - DTO for prompt config CREATE request body.
type PromptConfigCreateDTO struct { // skipcq: TCV-001
	Name                   string             `json:"name"            validate:"required"`
	ModelParameters        json.RawMessage    `json:"modelParameters" validate:"required"`
	ModelType              models.ModelType   `json:"modelType"       validate:"oneof=gpt-3.5-turbo gpt-3.5-turbo-16k gpt-4 gpt-4-32k"`
	ModelVendor            models.ModelVendor `json:"modelVendor"     validate:"oneof=OPEN_AI"`
	ProviderPromptMessages json.RawMessage    `json:"promptMessages"  validate:"required"`
	IsTest                 bool               `json:"isTest"`
}

// PromptConfigUpdateDTO - DTO for prompt config UPDATE request body.
type PromptConfigUpdateDTO struct { // skipcq: TCV-001
	Name                   *string             `json:"name,omitempty"            validate:"omitempty,required"`
	ModelParameters        *json.RawMessage    `json:"modelParameters,omitempty" validate:"omitempty,required"`
	ModelType              *models.ModelType   `json:"modelType,omitempty"       validate:"omitempty,oneof=gpt-3.5-turbo gpt-3.5-turbo-16k gpt-4 gpt-4-32k"`
	ModelVendor            *models.ModelVendor `json:"modelVendor,omitempty"     validate:"omitempty,oneof=OPEN_AI"`
	ProviderPromptMessages *json.RawMessage    `json:"promptMessages,omitempty"  validate:"omitempty,required"`
}

// ApplicationAPIKeyDTO - DTO for serializing application api key data.
type ApplicationAPIKeyDTO struct { // skipcq: TCV-001
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"           validate:"required"`
	Hash      *string   `json:"hash,omitempty"`
}

// AddUserAccountToProjectDTO - DTO for add user to project request body.
type AddUserAccountToProjectDTO struct { // skipcq: TCV-001
	UserID     string                      `json:"userId,omitempty" validate:"omitempty,required"`
	Email      string                      `json:"email,omitempty"  validate:"omitempty,required"`
	Permission models.AccessPermissionType `json:"permission"       validate:"required,oneof=ADMIN MEMBER"`
}

// UpdateUserAccountProjectPermissionDTO - DTO for update user account project permission request body.
type UpdateUserAccountProjectPermissionDTO struct { // skipcq: TCV-001
	UserID     string                      `json:"userId"     validate:"required"`
	Permission models.AccessPermissionType `json:"permission" validate:"required,oneof=ADMIN MEMBER"`
}

// UpdateUserAccountProjectPermissionDTO - DTO for serializing user account project + permission data.
type ProjectUserAccountDTO struct { // skipcq: TCV-001
	ID          string    `json:"id"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	FirebaseID  string    `json:"firebaseId"`
	PhoneNumber string    `json:"phoneNumber"`
	PhotoURL    string    `json:"photoUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	Permission  string    `json:"permission"`
}

// ApplicationAnalyticsDTO - DTO for serializing application analytics data.
type ApplicationAnalyticsDTO struct { // skipcq: TCV-001
	TotalRequests int64   `json:"totalRequests"`
	ProjectedCost float64 `json:"projectedCost"`
}

// ProjectAnalyticsDTO - DTO for serializing project analytics data.
type ProjectAnalyticsDTO struct { // skipcq: TCV-001
	TotalAPICalls int64   `json:"totalAPICalls"`
	ModelsCost    float64 `json:"modelsCost"`
}

// PromptConfigAnalyticsDTO - DTO for serializing prompt config analytics data.
type PromptConfigAnalyticsDTO struct { // skipcq: TCV-001
	TotalPromptRequests int64   `json:"totalPromptRequests"`
	ModelsCost          float64 `json:"modelsCost"`
}

// PromptConfigDTO - DTO for requesting a prompt config test.
type PromptConfigTestDTO struct { // skipcq: TCV-001
	Name                   string             `json:"name"                        validate:"required"`
	ModelParameters        json.RawMessage    `json:"modelParameters,omitempty"   validate:"omitempty,required"`
	ModelType              models.ModelType   `json:"modelType"                   validate:"oneof=gpt-3.5-turbo gpt-3.5-turbo-16k gpt-4 gpt-4-32k"`
	ModelVendor            models.ModelVendor `json:"modelVendor"                 validate:"oneof=OPEN_AI"`
	ProviderPromptMessages json.RawMessage    `json:"promptMessages,omitempty"    validate:"omitempty,required"`
	TemplateVariables      map[string]string  `json:"templateVariables,omitempty"`
	PromptConfigID         *string            `json:"promptConfigId,omitempty"    validate:"omitempty,required,uuid4"`
}

// PromptConfigTestResultDTO - DTO for serializing prompt config test results.
type PromptConfigTestResultDTO struct { // skipcq: TCV-001
	Content            *string `json:"content,omitempty"`
	ErrorMessage       *string `json:"errorMessage,omitempty"`
	FinishReason       *string `json:"finishReason,omitempty"`
	PromptConfigID     *string `json:"promptConfigId,omitempty"`
	PromptTestRecordID *string `json:"promptTestRecordId,omitempty"`
}

// OtpDTO - DTO for serializing an OTP.
type OtpDTO struct { // skipcq: TCV-001
	OTP string `json:"otp"`
}
