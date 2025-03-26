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
updated_at=$5,
image_url=$6
WHERE ID=$7
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
WHERE 
  ($1::TEXT = '' OR first_name ILIKE '%' || $1 || '%') AND
  ($2::TEXT = '' OR last_name ILIKE '%' || $2 || '%') AND
  ($3::TEXT = '' OR phone_no = $3) AND
  ($4::TEXT = '' OR email = $4)
ORDER BY first_name;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;