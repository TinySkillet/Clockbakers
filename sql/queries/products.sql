-- name: CreateProduct :one
INSERT INTO products(
 ID, SKU, name,
 description, price, stockQty,
 category
 ) VALUES (
 $1, $2, $3,
 $4, $5, $6,
 $7
 ) RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE SKU=$1;

-- name: UpdateProduct :one
UPDATE products SET 
SKU=$1, name=$2, description=$3,
price=$4, stockQty=$5, category=$6
WHERE SKU=$1
RETURNING *;

