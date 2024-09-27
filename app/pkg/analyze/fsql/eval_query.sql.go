// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: eval_query.sql

package fsql

import (
	"context"
	"time"
)

const getResult = `-- name: GetResult :one
SELECT id, time, proto_json FROM results
WHERE id = ?
`

func (q *Queries) GetResult(ctx context.Context, id string) (Result, error) {
	row := q.db.QueryRowContext(ctx, getResult, id)
	var i Result
	err := row.Scan(&i.ID, &i.Time, &i.ProtoJson)
	return i, err
}

const listResults = `-- name: ListResults :many
SELECT id, time, proto_json FROM results
WHERE (?1 = '' OR time < ?1)
ORDER BY time DESC
    LIMIT ?2
`

type ListResultsParams struct {
	Cursor   interface{}
	PageSize int64
}

// This queries for results.
// Results are listed in descending order of time (most recent first) because the primary use is for resuming
// in the evaluator
func (q *Queries) ListResults(ctx context.Context, arg ListResultsParams) ([]Result, error) {
	rows, err := q.db.QueryContext(ctx, listResults, arg.Cursor, arg.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Result
	for rows.Next() {
		var i Result
		if err := rows.Scan(&i.ID, &i.Time, &i.ProtoJson); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateResult = `-- name: UpdateResult :exec
INSERT OR REPLACE INTO results
(id, time, proto_json)
VALUES
(?, ?, ?)
`

type UpdateResultParams struct {
	ID        string
	Time      time.Time
	ProtoJson string
}

func (q *Queries) UpdateResult(ctx context.Context, arg UpdateResultParams) error {
	_, err := q.db.ExecContext(ctx, updateResult, arg.ID, arg.Time, arg.ProtoJson)
	return err
}
