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

type PopularProduct struct {
	ID           uuid.UUID `json:"id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	StockQty     int       `json:"stock_qty"`
	CategoryName string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	OrderCount   int       `json:"order_count"`
}

type CartProduct struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
}

func DBCartProductToCartProduct(dbCartProduct database.GetItemsFromCartRow) CartProduct {
	return CartProduct{
		SKU:         dbCartProduct.Sku,
		Name:        dbCartProduct.Name,
		Description: dbCartProduct.Description,
		Price:       dbCartProduct.Price,
		StockQty:    int(dbCartProduct.StockQty),
		Category:    dbCartProduct.Category,
		Quantity:    int(dbCartProduct.Quantity),
	}
}

func DBCartProductsToCartProducts(dbCartProducts []database.GetItemsFromCartRow) []CartProduct {
	cartProducts := make([]CartProduct, len(dbCartProducts))
	for i, dbCartProduct := range dbCartProducts {
		cartProducts[i] = DBCartProductToCartProduct(dbCartProduct)
	}
	return cartProducts
}

type Category struct {
	CategoryName string `json:"category_name"`
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

func DBPopularProductToPopularProduct(dbPopularProduct database.GetPopularItemsRow) PopularProduct {
	return PopularProduct{
		ID:           dbPopularProduct.ID,
		SKU:          dbPopularProduct.Sku,
		Name:         dbPopularProduct.Name,
		Description:  dbPopularProduct.Description,
		Price:        float64(dbPopularProduct.Price),
		StockQty:     int(dbPopularProduct.StockQty),
		CategoryName: dbPopularProduct.Category,
		CreatedAt:    dbPopularProduct.CreatedAt,
		UpdatedAt:    dbPopularProduct.UpdatedAt,
		OrderCount:   int(dbPopularProduct.OrderCount),
	}
}

func DBPopularProductsToPopularProducts(dbPopularProds []database.GetPopularItemsRow) []PopularProduct {
	popular_products := make([]PopularProduct, len(dbPopularProds))
	for i, dbPopularProd := range dbPopularProds {
		popular_products[i] = DBPopularProductToPopularProduct(dbPopularProd)
	}
	return popular_products
}
