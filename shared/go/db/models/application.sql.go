// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: application.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createApplication = `-- name: CreateApplication :one

INSERT INTO application (
    project_id,
    name,
    description
)
VALUES ($1, $2, $3)
RETURNING id, description, name, created_at, updated_at, deleted_at, project_id
`

type CreateApplicationParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ProjectID   pgtype.UUID `json:"projectId"`
}

// -- application
func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, createApplication, arg.ProjectID, arg.Name, arg.Description)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ProjectID,
	)
	return i, err
}

const deleteApplication = `-- name: DeleteApplication :exec
UPDATE application
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteApplication(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteApplication, id)
	return err
}

const retrieveApplication = `-- name: RetrieveApplication :one
SELECT
    id,
    description,
    name,
    created_at,
    updated_at,
    project_id
FROM application
WHERE
    id = $1
    AND deleted_at IS NULL
`

type RetrieveApplicationRow struct {
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
	Description string             `json:"description"`
	Name        string             `json:"name"`
	ID          pgtype.UUID        `json:"id"`
	ProjectID   pgtype.UUID        `json:"projectId"`
}

func (q *Queries) RetrieveApplication(ctx context.Context, id pgtype.UUID) (RetrieveApplicationRow, error) {
	row := q.db.QueryRow(ctx, retrieveApplication, id)
	var i RetrieveApplicationRow
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

const retrieveApplicationAPIRequestCount = `-- name: RetrieveApplicationAPIRequestCount :one
SELECT COUNT(prr.id) AS total_requests
FROM application AS a
INNER JOIN prompt_config AS pc ON a.id = pc.application_id
INNER JOIN prompt_request_record AS prr ON pc.id = prr.prompt_config_id
WHERE
    a.id = $1
    AND prr.created_at BETWEEN $2 AND $3
`

type RetrieveApplicationAPIRequestCountParams struct {
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	CreatedAt_2 pgtype.Timestamptz `json:"createdAt2"`
	ID          pgtype.UUID        `json:"id"`
}

func (q *Queries) RetrieveApplicationAPIRequestCount(ctx context.Context, arg RetrieveApplicationAPIRequestCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, retrieveApplicationAPIRequestCount, arg.ID, arg.CreatedAt, arg.CreatedAt_2)
	var total_requests int64
	err := row.Scan(&total_requests)
	return total_requests, err
}

const retrieveApplicationTokensTotalCost = `-- name: RetrieveApplicationTokensTotalCost :one
SELECT COALESCE(SUM(prr.request_tokens_cost + prr.response_tokens_cost), 0)
FROM application AS app
LEFT JOIN prompt_config AS pc ON app.id = pc.application_id
LEFT JOIN prompt_request_record AS prr ON pc.id = prr.prompt_config_id
WHERE
    app.id = $1
    AND prr.created_at BETWEEN $2 AND $3
`

type RetrieveApplicationTokensTotalCostParams struct {
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	CreatedAt_2 pgtype.Timestamptz `json:"createdAt2"`
	ID          pgtype.UUID        `json:"id"`
}

func (q *Queries) RetrieveApplicationTokensTotalCost(ctx context.Context, arg RetrieveApplicationTokensTotalCostParams) (pgtype.Numeric, error) {
	row := q.db.QueryRow(ctx, retrieveApplicationTokensTotalCost, arg.ID, arg.CreatedAt, arg.CreatedAt_2)
	var coalesce pgtype.Numeric
	err := row.Scan(&coalesce)
	return coalesce, err
}

const retrieveApplications = `-- name: RetrieveApplications :many
SELECT
    id,
    description,
    name,
    created_at,
    updated_at,
    project_id
FROM application
WHERE
    project_id = $1
    AND deleted_at IS NULL
`

type RetrieveApplicationsRow struct {
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
	Description string             `json:"description"`
	Name        string             `json:"name"`
	ID          pgtype.UUID        `json:"id"`
	ProjectID   pgtype.UUID        `json:"projectId"`
}

func (q *Queries) RetrieveApplications(ctx context.Context, projectID pgtype.UUID) ([]RetrieveApplicationsRow, error) {
	rows, err := q.db.Query(ctx, retrieveApplications, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RetrieveApplicationsRow
	for rows.Next() {
		var i RetrieveApplicationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
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
WHERE
    id = $1
    AND deleted_at IS NULL
RETURNING id, description, name, created_at, updated_at, deleted_at, project_id
`

type UpdateApplicationParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ID          pgtype.UUID `json:"id"`
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
		&i.DeletedAt,
		&i.ProjectID,
	)
	return i, err
}
