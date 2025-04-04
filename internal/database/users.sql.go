// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
ID, first_name, last_name,
email, phone_no,
address, password, role,
created_at, updated_at)
VALUES (
$1, $2, $3,
$4, $5,
$6, $7, $8,
$9, $10)
RETURNING id, first_name, last_name, email, phone_no, address, password, role, image_url, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	PhoneNo   sql.NullString
	Address   sql.NullString
	Password  string
	Role      UserType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNo,
		arg.Address,
		arg.Password,
		arg.Role,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Address,
		&i.Password,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, phone_no, address, password, role, image_url, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Address,
		&i.Password,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, email, phone_no, address, password, role, image_url, created_at, updated_at FROM users WHERE id=$1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Address,
		&i.Password,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name, last_name, email, phone_no, address, password, role, image_url, created_at, updated_at FROM users
WHERE 
  ($1::TEXT = '' OR first_name ILIKE '%' || $1 || '%') AND
  ($2::TEXT = '' OR last_name ILIKE '%' || $2 || '%') AND
  ($3::TEXT = '' OR phone_no = $3) AND
  ($4::TEXT = '' OR email = $4)
ORDER BY first_name
`

type GetUsersParams struct {
	Column1 string
	Column2 string
	Column3 string
	Column4 string
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNo,
			&i.Address,
			&i.Password,
			&i.Role,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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

const updateUser = `-- name: UpdateUser :one
UPDATE users SET 
first_name=$1,
last_name=$2,
phone_no=$3,
address=$4,
updated_at=$5,
image_url=$6
WHERE ID=$7
RETURNING id, first_name, last_name, email, phone_no, address, password, role, image_url, created_at, updated_at
`

type UpdateUserParams struct {
	FirstName string
	LastName  string
	PhoneNo   sql.NullString
	Address   sql.NullString
	UpdatedAt time.Time
	ImageUrl  sql.NullString
	ID        uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNo,
		arg.Address,
		arg.UpdatedAt,
		arg.ImageUrl,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Address,
		&i.Password,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
