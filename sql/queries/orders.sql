-- name: CreateOrder :one
INSERT INTO orders (
  ID, status, total_price,
  created_at, updated_at, user_id
) VALUES (
  $1, $2, $3,
  $4,
  $5,
  $6
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE ID = $1;

-- name: GetPopularItems :many
SELECT p.*, COUNT(oi.product_id) AS order_count
FROM products p
JOIN order_items oi ON p.ID = oi.product_id
GROUP BY p.ID
ORDER BY order_count DESC;

-- name: ListOrders :many
SELECT * FROM orders 
WHERE 
  ($1 IS NULL OR user_id = $1) AND 
  ($2 = '' OR status = $2)
ORDER BY created_at DESC;

-- name: UpdateOrderStatus :one
UPDATE orders SET 
  status = $2, updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE ID = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE ID = $1;