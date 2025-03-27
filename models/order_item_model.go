package models

import (
	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type OrderItem struct {
	ID              uuid.UUID `json:"id"`
	Quantity        int       `json:"quantity" validate:"required,gt=0"`
	Pounds          float64   `json:"pounds" validate:"required, gt=0"`
	Message         string    `json:"message"`
	PriceAtPurchase float64   `json:"price_at_purchase" validate:"required,gt=0"`
	OrderID         uuid.UUID `json:"order_id" validate:"required"`
	ProductID       uuid.UUID `json:"pid" validate:"required"`
}

func DBOrderItemToOrderItem(dbOrderItem database.OrderItem) OrderItem {
	return OrderItem{
		ID:              dbOrderItem.ID,
		Quantity:        int(dbOrderItem.Quantity),
		Pounds:          float64(dbOrderItem.Pounds),
		Message:         dbOrderItem.Message,
		PriceAtPurchase: float64(dbOrderItem.PriceAtPurchase),
		OrderID:         dbOrderItem.OrderID,
		ProductID:       dbOrderItem.ProductID,
	}
}

func DBOrderItemsToOrderItems(dbOrderItems []database.OrderItem) []OrderItem {
	orderItems := make([]OrderItem, len(dbOrderItems))
	for i, dbOrderItem := range dbOrderItems {
		orderItems[i] = DBOrderItemToOrderItem(dbOrderItem)
	}
	return orderItems
}
