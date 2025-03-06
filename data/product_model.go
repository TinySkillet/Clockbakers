package models

import (
	"regexp"

	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id"`
	SKU          string    `json:"sku" validate:"required,sku"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Price        float64   `json:"price" validate:"gt=0"`
	StockQty     int       `json:"stock_qty" validate:"required"`
	CategoryName string    `json:"category" validate:"required"`
}

type Category struct {
	CategoryName string `json:"name" validate:"required"`
}

func (p *Product) Validate() error {
	var Validate = validator.New()
	Validate.RegisterValidation("sku", validateSKU)
	return Validate.Struct(p)
}

func (c *Category) Validate() error {
	var validate = validator.New()
	return validate.Struct(c)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-ablkd
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	// if we don't have exactly one match
	if len(matches) != 1 {
		return false
	}
	return true
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
