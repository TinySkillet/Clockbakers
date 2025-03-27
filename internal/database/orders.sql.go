// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
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
) RETURNING id, status, total_price, quantity, pounds, message, delivery_time, delivery_date, created_at, updated_at, product_id, user_id
`

type CreateOrderParams struct {
	ID           uuid.UUID
	Status       OrderStatus
	TotalPrice   float32
	Quantity     int32
	Pounds       float32
	Message      string
	DeliveryTime DeliveryTimes
	DeliveryDate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductID    uuid.UUID
	UserID       uuid.UUID
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.Status,
		arg.TotalPrice,
		arg.Quantity,
		arg.Pounds,
		arg.Message,
		arg.DeliveryTime,
		arg.DeliveryDate,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ProductID,
		arg.UserID,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TotalPrice,
		&i.Quantity,
		&i.Pounds,
		&i.Message,
		&i.DeliveryTime,
		&i.DeliveryDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE ID = $1 AND user_id = $2
`

type DeleteOrderParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteOrder(ctx context.Context, arg DeleteOrderParams) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, arg.ID, arg.UserID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT o.id, o.status, o.total_price, o.quantity, o.pounds, o.message, o.delivery_time, o.delivery_date, o.created_at, o.updated_at, o.product_id, o.user_id, p.SKU, p.name, p.description, p.category 
FROM orders o
INNER JOIN products p ON
o.product_id = p.ID
WHERE o.ID = $1
`

type GetOrderRow struct {
	ID           uuid.UUID
	Status       OrderStatus
	TotalPrice   float32
	Quantity     int32
	Pounds       float32
	Message      string
	DeliveryTime DeliveryTimes
	DeliveryDate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductID    uuid.UUID
	UserID       uuid.UUID
	Sku          string
	Name         string
	Description  string
	Category     string
}

func (q *Queries) GetOrder(ctx context.Context, id uuid.UUID) (GetOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i GetOrderRow
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TotalPrice,
		&i.Quantity,
		&i.Pounds,
		&i.Message,
		&i.DeliveryTime,
		&i.DeliveryDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
		&i.Sku,
		&i.Name,
		&i.Description,
		&i.Category,
	)
	return i, err
}

const getPopularItems = `-- name: GetPopularItems :many
SELECT p.id, p.sku, p.name, p.description, p.price, p.stock_qty, p.category, p.created_at, p.updated_at, COUNT(o.product_id) AS order_count
FROM products p
JOIN orders o ON o.product_id = p.ID
GROUP BY p.ID
ORDER BY order_count DESC
`

type GetPopularItemsRow struct {
	ID          uuid.UUID
	Sku         string
	Name        string
	Description string
	Price       float32
	StockQty    int32
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OrderCount  int64
}

func (q *Queries) GetPopularItems(ctx context.Context) ([]GetPopularItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPopularItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPopularItemsRow
	for rows.Next() {
		var i GetPopularItemsRow
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
			&i.OrderCount,
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

const listOrders = `-- name: ListOrders :many
SELECT o.id, o.status, o.total_price, o.quantity, o.pounds, o.message, o.delivery_time, o.delivery_date, o.created_at, o.updated_at, o.product_id, o.user_id, p.SKU, p.name, p.description, p.category FROM orders o
INNER JOIN products p ON o.product_id = p.ID
WHERE 
  ($1::UUID IS NULL OR o.user_id = $1) AND 
  ($2::TEXT = '' OR o.status = $2)
ORDER BY o.created_at DESC
`

type ListOrdersParams struct {
	Column1 uuid.UUID
	Column2 string
}

type ListOrdersRow struct {
	ID           uuid.UUID
	Status       OrderStatus
	TotalPrice   float32
	Quantity     int32
	Pounds       float32
	Message      string
	DeliveryTime DeliveryTimes
	DeliveryDate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductID    uuid.UUID
	UserID       uuid.UUID
	Sku          string
	Name         string
	Description  string
	Category     string
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]ListOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListOrdersRow
	for rows.Next() {
		var i ListOrdersRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.TotalPrice,
			&i.Quantity,
			&i.Pounds,
			&i.Message,
			&i.DeliveryTime,
			&i.DeliveryDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductID,
			&i.UserID,
			&i.Sku,
			&i.Name,
			&i.Description,
			&i.Category,
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

const updateOrderStatus = `-- name: UpdateOrderStatus :one
UPDATE orders SET 
  status = $2, updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE ID = $1
RETURNING id, status, total_price, quantity, pounds, message, delivery_time, delivery_date, created_at, updated_at, product_id, user_id
`

type UpdateOrderStatusParams struct {
	ID     uuid.UUID
	Status OrderStatus
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrderStatus, arg.ID, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TotalPrice,
		&i.Quantity,
		&i.Pounds,
		&i.Message,
		&i.DeliveryTime,
		&i.DeliveryDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.UserID,
	)
	return i, err
}
