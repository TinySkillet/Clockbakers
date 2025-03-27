-- +goose Up
CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped',
'delivered', 'cancelled');

CREATE TYPE delivery_times AS ENUM ('morning', 'afternoon', 'evening');

CREATE TABLE orders (
  ID UUID PRIMARY KEY NOT NULL,
  status order_status NOT NULL DEFAULT 'pending',
  total_price float(10) NOT NULL,
  quantity INTEGER NOT NULL,
  pounds FLOAT(10) NOT NULL,
  message TEXT NOT NULL DEFAULT '',
  delivery_time delivery_times NOT NULL,
  delivery_date VARCHAR(64) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  product_id UUID NOT NULL REFERENCES products(ID),
  user_id UUID NOT NULL REFERENCES users(ID)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE orders;
DROP TYPE order_status;
DROP TYPE delivery_times;