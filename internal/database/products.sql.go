// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products(
 ID, SKU, name,
 description, price, stock_qty,
 category, created_at, updated_at
 ) VALUES (
 $1, $2, $3,
 $4, $5, $6,
 $7, $8, $9
 ) RETURNING id, sku, name, description, price, stock_qty, category, created_at, updated_at
`

type CreateProductParams struct {
	ID          uuid.UUID
	Sku         string
	Name        string
	Description string
	Price       float32
	StockQty    int32
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.Sku,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StockQty,
		arg.Category,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Sku,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQty,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE SKU=$1
`

func (q *Queries) DeleteProduct(ctx context.Context, sku string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, sku)
	return err
}

const getProducts = `-- name: GetProducts :many
SELECT id, sku, name, description, price, stock_qty, category, created_at, updated_at FROM products
WHERE 
  ($1::TEXT = '' OR name ILIKE '%' || $1 || '%') AND
  ($2 IS NULL OR price >= $2) AND
  ($3 IS NULL OR price <= $3) AND
  ($4::TEXT = '' OR category = $4)
ORDER BY name
`

type GetProductsParams struct {
	Column1 string
	Column2 interface{}
	Column3 interface{}
	Column4 string
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Sku,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.StockQty,
			&i.Category,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products SET 
SKU=$1, name=$2, description=$3,
price=$4, stock_qty=$5, category=$6
WHERE SKU=$1
RETURNING id, sku, name, description, price, stock_qty, category, created_at, updated_at
`

type UpdateProductParams struct {
	Sku         string
	Name        string
	Description string
	Price       float32
	StockQty    int32
	Category    string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Sku,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StockQty,
		arg.Category,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Sku,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQty,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
