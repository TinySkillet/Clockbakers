package models

import (
	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	Quantity  int32     `json:"quantity"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
}

func DBCartItemToCartItem(dbCartItem database.CartItem) CartItem {
	return CartItem{
		ID:        dbCartItem.ID,
		Quantity:  dbCartItem.Quantity,
		UserID:    dbCartItem.CartID,
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
