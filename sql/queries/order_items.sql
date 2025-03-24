-- name: CreateOrderItem :one
INSERT INTO order_items (
  ID, quantity, price_at_purchase,
  order_id, product_id)
VALUES(
  $1, $2, $3,
  $4, $5)
RETURNING *;

