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
UPDATE products
SET 
  SKU = COALESCE($1, SKU),
  name = COALESCE($2, name),
  description = COALESCE($3, description),
  price = COALESCE($4, price),
  stock_qty = COALESCE($5, stock_qty),
  category = COALESCE($6, category),
  updated_at = COALESCE($7, updated_at)
WHERE SKU = $1
RETURNING *;