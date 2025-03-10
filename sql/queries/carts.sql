-- name: CreateCart :one
INSERT INTO carts
(ID, user_id)
VALUES($1, $2)
RETURNING *;
