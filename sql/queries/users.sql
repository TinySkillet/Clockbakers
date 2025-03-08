-- name: CreateUser :one
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
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET 
first_name=$1,
last_name=$2,
phone_no=$3,
address=$4,
updated_at=$5
WHERE ID=$6
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
