// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE firebase_id = $1)
`

func (q *Queries) CheckUserExists(ctx context.Context, firebaseID string) (bool, error) {
	row := q.db.QueryRow(ctx, checkUserExists, firebaseID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createApplication = `-- name: CreateApplication :one
INSERT INTO application (
    project_id,
    name,
    description
)
VALUES ($1, $2, $3)
RETURNING id, description, name, created_at, updated_at, project_id
`

type CreateApplicationParams struct {
	ProjectID   pgtype.UUID `json:"projectId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, createApplication, arg.ProjectID, arg.Name, arg.Description)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO project (name, description)
VALUES ($1, $2)
RETURNING id, name, description, created_at
`

type CreateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject, arg.Name, arg.Description)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const createPromptConfig = `-- name: CreatePromptConfig :one
INSERT INTO prompt_config (
    name,
    model_parameters,
    model_type,
    model_vendor,
    provider_prompt_messages,
    expected_template_variables,
    is_default,
    application_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, model_parameters, model_type, model_vendor, provider_prompt_messages, expected_template_variables, is_default, created_at, updated_at, application_id
`

type CreatePromptConfigParams struct {
	Name                      string      `json:"name"`
	ModelParameters           []byte      `json:"modelParameters"`
	ModelType                 ModelType   `json:"modelType"`
	ModelVendor               ModelVendor `json:"modelVendor"`
	ProviderPromptMessages    []byte      `json:"providerPromptMessages"`
	ExpectedTemplateVariables []string    `json:"expectedTemplateVariables"`
	IsDefault                 bool        `json:"isDefault"`
	ApplicationID             pgtype.UUID `json:"applicationId"`
}

func (q *Queries) CreatePromptConfig(ctx context.Context, arg CreatePromptConfigParams) (PromptConfig, error) {
	row := q.db.QueryRow(ctx, createPromptConfig,
		arg.Name,
		arg.ModelParameters,
		arg.ModelType,
		arg.ModelVendor,
		arg.ProviderPromptMessages,
		arg.ExpectedTemplateVariables,
		arg.IsDefault,
		arg.ApplicationID,
	)
	var i PromptConfig
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ModelParameters,
		&i.ModelType,
		&i.ModelVendor,
		&i.ProviderPromptMessages,
		&i.ExpectedTemplateVariables,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApplicationID,
	)
	return i, err
}

const createPromptRequestRecord = `-- name: CreatePromptRequestRecord :one
INSERT INTO prompt_request_record (
    is_stream_response,
    request_tokens,
    response_tokens,
    start_time,
    finish_time,
    prompt_config_id,
    error_log
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, is_stream_response, request_tokens, response_tokens, start_time, finish_time, prompt_config_id, error_log
`

type CreatePromptRequestRecordParams struct {
	IsStreamResponse bool               `json:"isStreamResponse"`
	RequestTokens    int32              `json:"requestTokens"`
	ResponseTokens   int32              `json:"responseTokens"`
	StartTime        pgtype.Timestamptz `json:"startTime"`
	FinishTime       pgtype.Timestamptz `json:"finishTime"`
	PromptConfigID   pgtype.UUID        `json:"promptConfigId"`
	ErrorLog         pgtype.Text        `json:"errorLog"`
}

func (q *Queries) CreatePromptRequestRecord(ctx context.Context, arg CreatePromptRequestRecordParams) (PromptRequestRecord, error) {
	row := q.db.QueryRow(ctx, createPromptRequestRecord,
		arg.IsStreamResponse,
		arg.RequestTokens,
		arg.ResponseTokens,
		arg.StartTime,
		arg.FinishTime,
		arg.PromptConfigID,
		arg.ErrorLog,
	)
	var i PromptRequestRecord
	err := row.Scan(
		&i.ID,
		&i.IsStreamResponse,
		&i.RequestTokens,
		&i.ResponseTokens,
		&i.StartTime,
		&i.FinishTime,
		&i.PromptConfigID,
		&i.ErrorLog,
	)
	return i, err
}

const createPromptTest = `-- name: CreatePromptTest :one
INSERT INTO prompt_test (
    name,
    variable_values,
    response,
    created_at,
    prompt_request_record_id
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, variable_values, response, created_at, prompt_request_record_id
`

type CreatePromptTestParams struct {
	Name                  string             `json:"name"`
	VariableValues        []byte             `json:"variableValues"`
	Response              string             `json:"response"`
	CreatedAt             pgtype.Timestamptz `json:"createdAt"`
	PromptRequestRecordID pgtype.UUID        `json:"promptRequestRecordId"`
}

func (q *Queries) CreatePromptTest(ctx context.Context, arg CreatePromptTestParams) (PromptTest, error) {
	row := q.db.QueryRow(ctx, createPromptTest,
		arg.Name,
		arg.VariableValues,
		arg.Response,
		arg.CreatedAt,
		arg.PromptRequestRecordID,
	)
	var i PromptTest
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.VariableValues,
		&i.Response,
		&i.CreatedAt,
		&i.PromptRequestRecordID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (firebase_id)
VALUES ($1)
RETURNING id, firebase_id, created_at
`

func (q *Queries) CreateUser(ctx context.Context, firebaseID string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, firebaseID)
	var i User
	err := row.Scan(&i.ID, &i.FirebaseID, &i.CreatedAt)
	return i, err
}

const createUserProject = `-- name: CreateUserProject :one
INSERT INTO user_project (
    user_id, project_id, permission, is_user_default_project
)
VALUES ($1, $2, $3, $4)
RETURNING user_id, project_id, permission, is_user_default_project
`

type CreateUserProjectParams struct {
	UserID               pgtype.UUID          `json:"userId"`
	ProjectID            pgtype.UUID          `json:"projectId"`
	Permission           AccessPermissionType `json:"permission"`
	IsUserDefaultProject bool                 `json:"isUserDefaultProject"`
}

func (q *Queries) CreateUserProject(ctx context.Context, arg CreateUserProjectParams) (UserProject, error) {
	row := q.db.QueryRow(ctx, createUserProject,
		arg.UserID,
		arg.ProjectID,
		arg.Permission,
		arg.IsUserDefaultProject,
	)
	var i UserProject
	err := row.Scan(
		&i.UserID,
		&i.ProjectID,
		&i.Permission,
		&i.IsUserDefaultProject,
	)
	return i, err
}

const deleteApplication = `-- name: DeleteApplication :exec
DELETE
FROM application
WHERE id = $1
`

func (q *Queries) DeleteApplication(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteApplication, id)
	return err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE
FROM "project"
WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProject, id)
	return err
}

const deletePromptConfig = `-- name: DeletePromptConfig :exec
DELETE
FROM prompt_config
WHERE id = $1
`

func (q *Queries) DeletePromptConfig(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deletePromptConfig, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE firebase_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, firebaseID string) error {
	_, err := q.db.Exec(ctx, deleteUser, firebaseID)
	return err
}

const deleteUserProject = `-- name: DeleteUserProject :exec
DELETE
FROM "user_project"
WHERE project_id = $1
`

func (q *Queries) DeleteUserProject(ctx context.Context, projectID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserProject, projectID)
	return err
}

const findApplicationById = `-- name: FindApplicationById :one
SELECT id, description, name, created_at, updated_at, project_id -- noqa: L044
FROM application
WHERE id = $1
`

func (q *Queries) FindApplicationById(ctx context.Context, id pgtype.UUID) (Application, error) {
	row := q.db.QueryRow(ctx, findApplicationById, id)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}

const findDefaultPromptConfigByApplicationId = `-- name: FindDefaultPromptConfigByApplicationId :one
SELECT
    id,
    name,
    model_parameters,
    model_type,
    model_vendor,
    provider_prompt_messages,
    expected_template_variables,
    is_default,
    created_at,
    updated_at,
    application_id
FROM prompt_config
WHERE application_id = $1 AND is_default = TRUE
`

func (q *Queries) FindDefaultPromptConfigByApplicationId(ctx context.Context, applicationID pgtype.UUID) (PromptConfig, error) {
	row := q.db.QueryRow(ctx, findDefaultPromptConfigByApplicationId, applicationID)
	var i PromptConfig
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ModelParameters,
		&i.ModelType,
		&i.ModelVendor,
		&i.ProviderPromptMessages,
		&i.ExpectedTemplateVariables,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApplicationID,
	)
	return i, err
}

const findProjectsByUserId = `-- name: FindProjectsByUserId :many
SELECT
    p.created_at,
    p.description,
    p.id,
    p.name,
    up.is_user_default_project,
    up.permission
FROM project AS p
INNER JOIN user_project AS up ON p.id = up.project_id
WHERE up.user_id = $1
`

type FindProjectsByUserIdRow struct {
	CreatedAt            pgtype.Timestamptz   `json:"createdAt"`
	Description          string               `json:"description"`
	ID                   pgtype.UUID          `json:"id"`
	Name                 string               `json:"name"`
	IsUserDefaultProject bool                 `json:"isUserDefaultProject"`
	Permission           AccessPermissionType `json:"permission"`
}

func (q *Queries) FindProjectsByUserId(ctx context.Context, userID pgtype.UUID) ([]FindProjectsByUserIdRow, error) {
	rows, err := q.db.Query(ctx, findProjectsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindProjectsByUserIdRow
	for rows.Next() {
		var i FindProjectsByUserIdRow
		if err := rows.Scan(
			&i.CreatedAt,
			&i.Description,
			&i.ID,
			&i.Name,
			&i.IsUserDefaultProject,
			&i.Permission,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPromptConfigById = `-- name: FindPromptConfigById :one
SELECT
    id,
    name,
    model_parameters,
    model_type,
    model_vendor,
    provider_prompt_messages,
    expected_template_variables,
    is_default,
    created_at,
    updated_at,
    application_id
FROM prompt_config
WHERE id = $1
`

func (q *Queries) FindPromptConfigById(ctx context.Context, id pgtype.UUID) (PromptConfig, error) {
	row := q.db.QueryRow(ctx, findPromptConfigById, id)
	var i PromptConfig
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ModelParameters,
		&i.ModelType,
		&i.ModelVendor,
		&i.ProviderPromptMessages,
		&i.ExpectedTemplateVariables,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApplicationID,
	)
	return i, err
}

const findPromptRequestRecords = `-- name: FindPromptRequestRecords :many
SELECT
    id,
    is_stream_response,
    request_tokens,
    response_tokens,
    start_time,
    finish_time,
    prompt_config_id,
    error_log
FROM prompt_request_record
WHERE prompt_config_id = $1
ORDER BY start_time DESC
`

func (q *Queries) FindPromptRequestRecords(ctx context.Context, promptConfigID pgtype.UUID) ([]PromptRequestRecord, error) {
	rows, err := q.db.Query(ctx, findPromptRequestRecords, promptConfigID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PromptRequestRecord
	for rows.Next() {
		var i PromptRequestRecord
		if err := rows.Scan(
			&i.ID,
			&i.IsStreamResponse,
			&i.RequestTokens,
			&i.ResponseTokens,
			&i.StartTime,
			&i.FinishTime,
			&i.PromptConfigID,
			&i.ErrorLog,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPromptTests = `-- name: FindPromptTests :many
SELECT
    prc.finish_time,
    prc.request_tokens,
    prc.start_time,
    prc.is_stream_response,
    pt.created_at,
    pt.id,
    pt.name,
    pt.response,
    pt.variable_values
FROM prompt_test AS pt
LEFT JOIN prompt_request_record AS prc
    ON pt.prompt_request_record_id = prc.id
WHERE prc.prompt_config_id = $1
ORDER BY pt.created_at
`

type FindPromptTestsRow struct {
	FinishTime       pgtype.Timestamptz `json:"finishTime"`
	RequestTokens    pgtype.Int4        `json:"requestTokens"`
	StartTime        pgtype.Timestamptz `json:"startTime"`
	IsStreamResponse pgtype.Bool        `json:"isStreamResponse"`
	CreatedAt        pgtype.Timestamptz `json:"createdAt"`
	ID               pgtype.UUID        `json:"id"`
	Name             string             `json:"name"`
	Response         string             `json:"response"`
	VariableValues   []byte             `json:"variableValues"`
}

func (q *Queries) FindPromptTests(ctx context.Context, promptConfigID pgtype.UUID) ([]FindPromptTestsRow, error) {
	rows, err := q.db.Query(ctx, findPromptTests, promptConfigID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindPromptTestsRow
	for rows.Next() {
		var i FindPromptTestsRow
		if err := rows.Scan(
			&i.FinishTime,
			&i.RequestTokens,
			&i.StartTime,
			&i.IsStreamResponse,
			&i.CreatedAt,
			&i.ID,
			&i.Name,
			&i.Response,
			&i.VariableValues,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserByFirebaseId = `-- name: FindUserByFirebaseId :one
SELECT
    id,
    firebase_id,
    created_at
FROM "user"
WHERE firebase_id = $1
`

func (q *Queries) FindUserByFirebaseId(ctx context.Context, firebaseID string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByFirebaseId, firebaseID)
	var i User
	err := row.Scan(&i.ID, &i.FirebaseID, &i.CreatedAt)
	return i, err
}

const retrieveApplicationPromptConfigs = `-- name: RetrieveApplicationPromptConfigs :many
SELECT
    id,
    name,
    model_parameters,
    model_type,
    model_vendor,
    provider_prompt_messages,
    expected_template_variables,
    is_default,
    created_at,
    updated_at,
    application_id
FROM prompt_config
WHERE application_id = $1
`

func (q *Queries) RetrieveApplicationPromptConfigs(ctx context.Context, applicationID pgtype.UUID) ([]PromptConfig, error) {
	rows, err := q.db.Query(ctx, retrieveApplicationPromptConfigs, applicationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PromptConfig
	for rows.Next() {
		var i PromptConfig
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ModelParameters,
			&i.ModelType,
			&i.ModelVendor,
			&i.ProviderPromptMessages,
			&i.ExpectedTemplateVariables,
			&i.IsDefault,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ApplicationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateApplication = `-- name: UpdateApplication :one
UPDATE application
SET
    name = $2,
    description = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING id, description, name, created_at, updated_at, project_id
`

type UpdateApplicationParams struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

func (q *Queries) UpdateApplication(ctx context.Context, arg UpdateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, updateApplication, arg.ID, arg.Name, arg.Description)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}

const updatePromptConfig = `-- name: UpdatePromptConfig :one
UPDATE prompt_config
SET
    name = $2,
    model_parameters = $3,
    model_type = $4,
    is_default = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, model_parameters, model_type, model_vendor, provider_prompt_messages, expected_template_variables, is_default, created_at, updated_at, application_id
`

type UpdatePromptConfigParams struct {
	ID              pgtype.UUID `json:"id"`
	Name            string      `json:"name"`
	ModelParameters []byte      `json:"modelParameters"`
	ModelType       ModelType   `json:"modelType"`
	IsDefault       bool        `json:"isDefault"`
}

func (q *Queries) UpdatePromptConfig(ctx context.Context, arg UpdatePromptConfigParams) (PromptConfig, error) {
	row := q.db.QueryRow(ctx, updatePromptConfig,
		arg.ID,
		arg.Name,
		arg.ModelParameters,
		arg.ModelType,
		arg.IsDefault,
	)
	var i PromptConfig
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ModelParameters,
		&i.ModelType,
		&i.ModelVendor,
		&i.ProviderPromptMessages,
		&i.ExpectedTemplateVariables,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApplicationID,
	)
	return i, err
}

const updatePromptConfigIsDefault = `-- name: UpdatePromptConfigIsDefault :exec
UPDATE prompt_config
SET
    is_default = $2,
    updated_at = NOW()
WHERE id = $1
`

type UpdatePromptConfigIsDefaultParams struct {
	ID        pgtype.UUID `json:"id"`
	IsDefault bool        `json:"isDefault"`
}

func (q *Queries) UpdatePromptConfigIsDefault(ctx context.Context, arg UpdatePromptConfigIsDefaultParams) error {
	_, err := q.db.Exec(ctx, updatePromptConfigIsDefault, arg.ID, arg.IsDefault)
	return err
}
