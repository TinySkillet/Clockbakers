-- +goose Up
CREATE TABLE order_items (
  ID UUID NOT NULL PRIMARY KEY,
  quantity INTEGER NOT NULL,
  pounds FLOAT(10) NOT NULL,
  message TEXT NOT NULL DEFAULT '',
  price_at_purchase FLOAT(10) NOT NULL,
  order_id UUID NOT NULL REFERENCES orders(ID),
  product_id UUID NOT NULL REFERENCES products(ID)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE order_items;