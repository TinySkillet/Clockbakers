-- name: CreateDeliveryAddress :one
INSERT INTO delivery_addresses (
  ID, address, user_id
)
VALUES (
$1, $2, $3
) RETURNING *;

-- name: GetDeliveryAddresses :many
SELECT * FROM delivery_addresses
WHERE user_id = $1;

-- name: DeleteDeliveryAddress :exec
DELETE FROM delivery_addresses
WHERE ID = $1 AND user_id = $2;