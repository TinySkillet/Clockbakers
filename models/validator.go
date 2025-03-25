package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// product validation
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// ensures SKU follows "abc-def-ghi" pattern
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-z]+-[a-z]+-[a-z]+$`)
	return re.MatchString(fl.Field().String())
}

// review validation
func (r *Review) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// user validation
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// order validation
func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

// orderitem validation
func (oi *OrderItem) Validate() error {
	validate := validator.New()
	return validate.Struct(oi)
}

// cartitem validation
func (ci *CartItem) Validate() error {
	validate := validator.New()
	return validate.Struct(ci)
}

// login request validation
func (lq *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(lq)
}

func (c *Category) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
