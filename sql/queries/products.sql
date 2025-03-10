-- name: CreateProduct :one
INSERT INTO products(
 ID, SKU, name,
 description, price, stock_qty,
 category, created_at, updated_at
 ) VALUES (
 $1, $2, $3,
 $4, $5, $6,
 $7, $8, $9
 ) RETURNING *;

-- name: GetProducts :many
SELECT * FROM products
WHERE 
  ($1::TEXT = '' OR name ILIKE '%' || $1 || '%') AND
  ($2::NUMERIC IS NULL OR price >= $2) AND
  ($3::NUMERIC IS NULL OR price <= $3) AND
  ($4::TEXT = '' OR category = $4)
ORDER BY name;


-- name: DeleteProduct :exec
DELETE FROM products WHERE SKU=$1;

-- name: UpdateProduct :one
UPDATE products SET 
SKU=$1, name=$2, description=$3,
price=$4, stock_qty=$5, category=$6,
updated_at=$7
WHERE SKU=$1
RETURNING *;