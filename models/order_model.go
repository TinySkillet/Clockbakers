package models

import (
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type ListOrderRow struct {
	ID           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	TotalPrice   float32   `json:"total_price" validate:"required,gt=0"`
	Quantity     int       `json:"quantity" validate:"required,gt=0"`
	Pounds       float64   `json:"pounds" validate:"required"`
	Message      string    `json:"message"`
	DeliveryTime string    `json:"delivery_time" validate:"required"`
	DeliveryDate string    `json:"delivery_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	ProductID    uuid.UUID `json:"product_id" validate:"required"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
}

type Order struct {
	ID           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	TotalPrice   float32   `json:"total_price" validate:"required,gt=0"`
	Quantity     int       `json:"quantity" validate:"required,gt=0"`
	Pounds       float64   `json:"pounds" validate:"required"`
	Message      string    `json:"message"`
	DeliveryTime string    `json:"delivery_time" validate:"required"`
	DeliveryDate string    `json:"delivery_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	ProductID    uuid.UUID `json:"product_id" validate:"required"`
}

func DBListOrderRowToOrderRow(dbOrderRow database.ListOrdersRow) ListOrderRow {
	return ListOrderRow{
		ID:           dbOrderRow.ID,
		Status:       string(dbOrderRow.Status),
		TotalPrice:   dbOrderRow.TotalPrice,
		Quantity:     int(dbOrderRow.Quantity),
		Pounds:       float64(dbOrderRow.Pounds),
		Message:      dbOrderRow.Message,
		DeliveryTime: string(dbOrderRow.DeliveryTime),
		CreatedAt:    dbOrderRow.CreatedAt,
		UpdatedAt:    dbOrderRow.UpdatedAt,
		UserID:       dbOrderRow.UserID,
		ProductID:    dbOrderRow.ProductID,
		SKU:          dbOrderRow.Sku,
		Name:         dbOrderRow.Name,
		Description:  dbOrderRow.Description,
		Category:     dbOrderRow.Category,
	}
}

func DBListOrderRowsToOrderRows(dbOrderRows []database.ListOrdersRow) []ListOrderRow {
	orderRows := make([]ListOrderRow, len(dbOrderRows))
	for i, dbOrderRow := range dbOrderRows {
		orderRows[i] = DBListOrderRowToOrderRow(dbOrderRow)
	}
	return orderRows
}

func DBOrderToOrder(dbOrder database.Order) Order {
	return Order{
		ID:           dbOrder.ID,
		Status:       string(dbOrder.Status),
		DeliveryTime: string(dbOrder.DeliveryTime),
		TotalPrice:   dbOrder.TotalPrice,
		CreatedAt:    dbOrder.CreatedAt,
		UpdatedAt:    dbOrder.UpdatedAt,
		UserID:       dbOrder.UserID,
	}
}

func DBOrdersToOrders(dbOrders []database.Order) []Order {
	orders := make([]Order, len(dbOrders))
	for i, dbOrder := range dbOrders {
		orders[i] = DBOrderToOrder(dbOrder)
	}
	return orders
}
