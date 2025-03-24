package models

import (
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `json:"id"`
	Status     string
	TotalPrice float32     `json:"total_price"`
	Items      []OrderItem `json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	UserID     uuid.UUID   `json:"user_id"`
}

func DBOrderToOrder(dbOrder database.Order) Order {
	return Order{
		ID:         dbOrder.ID,
		Status:     string(dbOrder.Status),
		TotalPrice: dbOrder.TotalPrice,
		CreatedAt:  dbOrder.CreatedAt,
		UpdatedAt:  dbOrder.UpdatedAt,
		UserID:     dbOrder.UserID,
	}
}

func DBOrdersToOrders(dbOrders []database.Order) []Order {
	orders := make([]Order, len(dbOrders))
	for i, dbOrder := range dbOrders {
		orders[i] = DBOrderToOrder(dbOrder)
	}
	return orders
}
