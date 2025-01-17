// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const insertValues = `-- name: InsertValues :one
INSERT INTO
    test (name)
VALUES
    ('one') RETURNING id, name, created_at, updated_at
`

func (q *Queries) InsertValues(ctx context.Context) (Test, error) {
	row := q.db.QueryRow(ctx, insertValues)
	var i Test
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
