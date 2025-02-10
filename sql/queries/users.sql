-- name: CreateUser :one
INSERT INTO users (
ID, first_name, last_name,
email, phone_no,
address, password, role
)
VALUES (
$1, $2, $3,
$4, $5,
$6, $7, $8)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUsersByName :many
SELECT * FROM users WHERE first_name || ' ' || last_name LIKE $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;

-- name: GetRoleByIdAndEmail :one
SELECT role FROM users WHERE id=$1 AND email=$2;
