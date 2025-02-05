// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: insert_contact_request.sql

package database

import (
	"context"
	"database/sql"
)

const insertContactRequest = `-- name: InsertContactRequest :one
INSERT INTO contact_requests (name, email, text, telephone)
VALUES (?, ?, ?, ?)
RETURNING id, name, email, text, telephone
`

type InsertContactRequestParams struct {
	Name      string
	Email     string
	Text      sql.NullString
	Telephone sql.NullString
}

func (q *Queries) InsertContactRequest(ctx context.Context, arg InsertContactRequestParams) (ContactRequest, error) {
	row := q.db.QueryRowContext(ctx, insertContactRequest,
		arg.Name,
		arg.Email,
		arg.Text,
		arg.Telephone,
	)
	var i ContactRequest
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Text,
		&i.Telephone,
	)
	return i, err
}
