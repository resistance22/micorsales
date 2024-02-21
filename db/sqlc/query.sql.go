// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, email, password, created_at, first_name, last_name FROM app_user
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (AppUser, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i AppUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}