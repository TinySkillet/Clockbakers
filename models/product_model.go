package models

import (
	"time"

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
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
		StockQty:     int(dbProd.StockQty),
		CategoryName: dbProd.Category,
		CreatedAt:    dbProd.CreatedAt,
		UpdatedAt:    dbProd.UpdatedAt,
	}
}

func DBProductsToProducts(dbProds []database.Product) []Product {
	products := make([]Product, len(dbProds))
	for i, dbProd := range dbProds {
		products[i] = DBProductToProduct(dbProd)
	}
	return products
}
