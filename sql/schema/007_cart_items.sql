-- +goose Up
CREATE TABLE cart_items (
  ID UUID PRIMARY KEY NOT NULL,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL REFERENCES carts(ID)
  ON DELETE CASCADE,
  product_id UUID NOT NULL REFERENCES products(ID)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE cart_items;