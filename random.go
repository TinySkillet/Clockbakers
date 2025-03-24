package main

import (
	"database/sql"
	"time"
)

// need to update this
type UpdateProductParams struct {
	Sku         sql.NullString  `json:"sku"`
	Name        sql.NullString  `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       sql.NullFloat64 `json:"price"`
	StockQty    sql.NullInt32   `json:"stock_qty"`
	Category    sql.NullString  `json:"category"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
