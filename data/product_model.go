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
	Price        float64   `json:"price" validate:"required,gt=0"`
	StockQty     int       `json:"stock_qty" validate:"required,gte=0"`
	CategoryName string    `json:"category" validate:"required"`
}

type Category struct {
	CategoryName string `json:"name" validate:"required"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func (c *Category) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

// ensures SKU follows "abc-def-ghi" pattern
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-z]+-[a-z]+-[a-z]+$`)
	return re.MatchString(fl.Field().String())
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
	}
}
