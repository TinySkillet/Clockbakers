-- +goose Up
CREATE TABLE products(
  ID UUID PRIMARY KEY NOT NULL,
  SKU VARCHAR(32) UNIQUE NOT NULL,
  name VARCHAR(64) NOT NULL,
  description VARCHAR(64) NOT NULL,
  price float(10) NOT NULL,
  stock_qty INTEGER NOT NULL,
  category VARCHAR(32) NOT NULL REFERENCES categories(name)
  ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE products;