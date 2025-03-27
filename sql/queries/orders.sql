-- name: CreateOrder :one
INSERT INTO orders (
  ID, status, total_price, quantity,
  pounds, message,
  delivery_time, delivery_date, created_at, updated_at,
  product_id, user_id
) VALUES (
  $1, $2, $3, $4,
  $5, $6,
  $7, $8, $9, $10,
  $11, $12
) RETURNING *;

-- name: GetOrder :one
SELECT o.*, p.SKU, p.name, p.description, p.category 
FROM orders o
INNER JOIN products p ON
o.product_id = p.ID
WHERE o.ID = $1;

-- name: GetPopularItems :many
SELECT p.*, COUNT(o.product_id) AS order_count
FROM products p
JOIN orders o ON o.product_id = p.ID
GROUP BY p.ID
ORDER BY order_count DESC;

-- name: ListOrders :many
SELECT o.*, p.SKU, p.name, p.description, p.category FROM orders o
INNER JOIN products p ON o.product_id = p.ID
WHERE 
  ($1::UUID IS NULL OR user_id = $1) AND 
  ($2::TEXT = '' OR status = $2)
ORDER BY o.created_at DESC;

-- name: UpdateOrderStatus :one
UPDATE orders SET 
  status = $2, updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE ID = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE ID = $1 AND user_id = $2;