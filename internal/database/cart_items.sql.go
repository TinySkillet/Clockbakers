// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cart_items.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addToCart = `-- name: AddToCart :exec
INSERT INTO cart_items (
  ID, quantity,
  cart_id, product_id
) VALUES (
  $1, $2,
  $3, $4
)
`

type AddToCartParams struct {
	ID        uuid.UUID
	Quantity  int32
	CartID    uuid.UUID
	ProductID uuid.UUID
}

func (q *Queries) AddToCart(ctx context.Context, arg AddToCartParams) error {
	_, err := q.db.ExecContext(ctx, addToCart,
		arg.ID,
		arg.Quantity,
		arg.CartID,
		arg.ProductID,
	)
	return err
}

const getItemsFromCart = `-- name: GetItemsFromCart :many
SELECT p.SKU, p.name, p.description, p.price,
       p.stock_qty, p.category, c.quantity
FROM products p
INNER JOIN cart_items c ON  p.ID = c.product_id
WHERE c.cart_id = $1
`

type GetItemsFromCartRow struct {
	Sku         string
	Name        string
	Description string
	Price       float32
	StockQty    int32
	Category    string
	Quantity    int32
}

func (q *Queries) GetItemsFromCart(ctx context.Context, cartID uuid.UUID) ([]GetItemsFromCartRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsFromCart, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsFromCartRow
	for rows.Next() {
		var i GetItemsFromCartRow
		if err := rows.Scan(
			&i.Sku,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.StockQty,
			&i.Category,
			&i.Quantity,
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

const reduceQuantityFromCart = `-- name: ReduceQuantityFromCart :exec
UPDATE cart_items
SET quantity = quantity - $1
WHERE product_id = $2 
AND cart_id = $3
AND quantity > $1
`

type ReduceQuantityFromCartParams struct {
	Quantity  int32
	ProductID uuid.UUID
	CartID    uuid.UUID
}

func (q *Queries) ReduceQuantityFromCart(ctx context.Context, arg ReduceQuantityFromCartParams) error {
	_, err := q.db.ExecContext(ctx, reduceQuantityFromCart, arg.Quantity, arg.ProductID, arg.CartID)
	return err
}

const removeFromCart = `-- name: RemoveFromCart :exec
DELETE FROM cart_items 
WHERE product_id = $1 
AND cart_id = $2
`

type RemoveFromCartParams struct {
	ProductID uuid.UUID
	CartID    uuid.UUID
}

func (q *Queries) RemoveFromCart(ctx context.Context, arg RemoveFromCartParams) error {
	_, err := q.db.ExecContext(ctx, removeFromCart, arg.ProductID, arg.CartID)
	return err
}
