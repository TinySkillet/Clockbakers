-- name: CreateCategory :one
INSERT INTO categories (
name
)
VALUES ($1)
RETURNING *;

-- name: UpdateCategory :one
UPDATE categories SET name=$1 WHERE name=$2
RETURNING *;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE name=$1;

