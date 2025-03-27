-- name: AddToCart :exec
INSERT INTO cart_items (
  ID, quantity,
  cart_id, product_id
) VALUES (
  $1, $2,
  $3, $4
);

-- name: ReduceQuantityFromCart :exec
UPDATE cart_items
SET quantity = quantity - $1
WHERE product_id = $2 
AND cart_id = $3
AND quantity > $1;


-- name: RemoveFromCart :exec
DELETE FROM cart_items 
WHERE product_id = $1 
AND cart_id = $2;


-- name: GetItemsFromCart :many
SELECT p.SKU, p.name, p.description, p.price,
       p.stock_qty, p.category, c.quantity
FROM products p
INNER JOIN cart_items c ON  p.ID = c.product_id
WHERE c.cart_id = $1;
