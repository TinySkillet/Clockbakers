-- name: CreateCart :one
INSERT INTO carts
(ID, user_id)
VALUES($1, $2)
RETURNING *;


-- name: GetCartID :one
SELECT ID FROM carts WHERE user_id=$1;