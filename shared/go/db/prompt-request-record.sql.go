// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: prompt-request-record.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPromptRequestRecord = `-- name: CreatePromptRequestRecord :one

INSERT INTO prompt_request_record (
    is_stream_response,
    request_tokens,
    response_tokens,
    start_time,
    finish_time,
    stream_response_latency,
    prompt_config_id,
    error_log
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, is_stream_response, request_tokens, response_tokens, start_time, finish_time, stream_response_latency, prompt_config_id, error_log, created_at, deleted_at
`

type CreatePromptRequestRecordParams struct {
	IsStreamResponse      bool               `json:"isStreamResponse"`
	RequestTokens         int32              `json:"requestTokens"`
	ResponseTokens        int32              `json:"responseTokens"`
	StartTime             pgtype.Timestamptz `json:"startTime"`
	FinishTime            pgtype.Timestamptz `json:"finishTime"`
	StreamResponseLatency pgtype.Int8        `json:"streamResponseLatency"`
	PromptConfigID        pgtype.UUID        `json:"promptConfigId"`
	ErrorLog              pgtype.Text        `json:"errorLog"`
}

// -- prompt request record
func (q *Queries) CreatePromptRequestRecord(ctx context.Context, arg CreatePromptRequestRecordParams) (PromptRequestRecord, error) {
	row := q.db.QueryRow(ctx, createPromptRequestRecord,
		arg.IsStreamResponse,
		arg.RequestTokens,
		arg.ResponseTokens,
		arg.StartTime,
		arg.FinishTime,
		arg.StreamResponseLatency,
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
		&i.StreamResponseLatency,
		&i.PromptConfigID,
		&i.ErrorLog,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const retrieveTotalPromptRequestRecord = `-- name: RetrieveTotalPromptRequestRecord :one
SELECT COUNT(prr.id) AS total_requests
FROM prompt_request_record AS prr
INNER JOIN prompt_config AS pc ON prr.prompt_config_id = pc.id
WHERE
    pc.application_id = $1
    AND prr.created_at BETWEEN $2 AND $3
`

type RetrieveTotalPromptRequestRecordParams struct {
	ApplicationID pgtype.UUID        `json:"applicationId"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	CreatedAt_2   pgtype.Timestamptz `json:"createdAt2"`
}

func (q *Queries) RetrieveTotalPromptRequestRecord(ctx context.Context, arg RetrieveTotalPromptRequestRecordParams) (int64, error) {
	row := q.db.QueryRow(ctx, retrieveTotalPromptRequestRecord, arg.ApplicationID, arg.CreatedAt, arg.CreatedAt_2)
	var total_requests int64
	err := row.Scan(&total_requests)
	return total_requests, err
}

const retrieveTotalTokensPerPromptConfig = `-- name: RetrieveTotalTokensPerPromptConfig :many
SELECT
    pc.model_type,
    SUM(prr.request_tokens + prr.response_tokens) AS total_tokens
FROM prompt_request_record AS prr
INNER JOIN prompt_config AS pc ON prr.prompt_config_id = pc.id
WHERE
    pc.application_id = $1
    AND prr.created_at BETWEEN $2 AND $3
GROUP BY pc.model_type
`

type RetrieveTotalTokensPerPromptConfigParams struct {
	ApplicationID pgtype.UUID        `json:"applicationId"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	CreatedAt_2   pgtype.Timestamptz `json:"createdAt2"`
}

type RetrieveTotalTokensPerPromptConfigRow struct {
	ModelType   ModelType `json:"modelType"`
	TotalTokens int64     `json:"totalTokens"`
}

func (q *Queries) RetrieveTotalTokensPerPromptConfig(ctx context.Context, arg RetrieveTotalTokensPerPromptConfigParams) ([]RetrieveTotalTokensPerPromptConfigRow, error) {
	rows, err := q.db.Query(ctx, retrieveTotalTokensPerPromptConfig, arg.ApplicationID, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RetrieveTotalTokensPerPromptConfigRow
	for rows.Next() {
		var i RetrieveTotalTokensPerPromptConfigRow
		if err := rows.Scan(&i.ModelType, &i.TotalTokens); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
