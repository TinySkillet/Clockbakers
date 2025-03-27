package models

import (
	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	Quantity  int32     `json:"quantity" validate:"required,gt=0"`
	CartID    uuid.UUID `json:"cart_id" validate:"required"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
}

type CartItemResponse struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int32   `json:"stock_quantity"`
	Category    string  `json:"category"`
	Quantity    int32   `json:"cart_quantity"`
}

func DBCartItemToCartItem(dbCartItem database.CartItem) CartItem {
	return CartItem{
		ID:        dbCartItem.ID,
		Quantity:  dbCartItem.Quantity,
		CartID:    dbCartItem.CartID,
		ProductID: dbCartItem.ProductID,
	}
}

func DBCartItemsToCartITems(dbCartItems []database.CartItem) []CartItem {
	cartItems := make([]CartItem, len(dbCartItems))
	for i, dbCartItem := range dbCartItems {
		cartItems[i] = DBCartItemToCartItem(dbCartItem)
	}
	return cartItems
}
