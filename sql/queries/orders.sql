-- name: CreateOrder :one
INSERT INTO orders (
  ID, status, total_price,
  created_at, updated_at, user_id
) VALUES (
  $1, $2, $3,
  CURRENT_TIMESTAMP AT TIME ZONE 'UTC',
  CURRENT_TIMESTAMP AT TIME ZONE 'UTC',
  $4
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE ID = $1;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY created_at DESC;

-- name: GetOrdersByUserID :many
SELECT * FROM orders WHERE user_id = $1;

-- name: UpdateOrderStatus :one
UPDATE orders SET 
  status = $2, updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE ID = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE ID = $1;
