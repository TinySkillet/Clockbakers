package models

import (
	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	StockQty     int       `json:"stock_qty"`
	CategoryName string    `json:"category"`
}

type Category struct {
	CategoryName string `json:"name"`
}

func DBProductToProduct(dbProd database.Product) Product {
	return Product{
		ID:           dbProd.ID,
		SKU:          dbProd.Sku,
		Name:         dbProd.Name,
		Description:  dbProd.Description,
		Price:        float64(dbProd.Price),
		StockQty:     int(dbProd.Stockqty),
		CategoryName: dbProd.Category,
	}
}
