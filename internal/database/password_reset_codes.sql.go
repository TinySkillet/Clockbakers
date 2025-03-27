// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: password_reset_codes.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPasswordResetCode = `-- name: CreatePasswordResetCode :one
INSERT INTO password_reset_codes (id, email, code, expires_at, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreatePasswordResetCodeParams struct {
	ID        uuid.UUID
	Email     string
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (q *Queries) CreatePasswordResetCode(ctx context.Context, arg CreatePasswordResetCodeParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createPasswordResetCode,
		arg.ID,
		arg.Email,
		arg.Code,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteExpiredPasswordResetCodes = `-- name: DeleteExpiredPasswordResetCodes :exec
DELETE FROM password_reset_codes
WHERE expires_at <= CURRENT_TIMESTAMP
`

func (q *Queries) DeleteExpiredPasswordResetCodes(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteExpiredPasswordResetCodes)
	return err
}

const getValidPasswordResetCode = `-- name: GetValidPasswordResetCode :one
SELECT id, email, code, expires_at, created_at FROM password_reset_codes
WHERE email = $1 AND expires_at > CURRENT_TIMESTAMP
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetValidPasswordResetCode(ctx context.Context, email string) (PasswordResetCode, error) {
	row := q.db.QueryRowContext(ctx, getValidPasswordResetCode, email)
	var i PasswordResetCode
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Code,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
