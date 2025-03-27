-- name: CreateOrderItem :one
INSERT INTO order_items (
  ID, quantity, pounds, message, price_at_purchase,
  order_id, product_id)
VALUES(
  $1, $2, $3,
  $4, $5, $6, $7)
RETURNING *;

-- name: GetOrderItemsByOrderID :many
SELECT * FROM order_items 
WHERE order_id = $1;
