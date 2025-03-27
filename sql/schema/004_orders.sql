-- +goose Up
CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped',
'delivered', 'cancelled');

CREATE TYPE delivery_times AS ENUM ('morning', 'afternoon', 'evening');

CREATE TABLE orders (
  ID UUID PRIMARY KEY NOT NULL,
  status order_status NOT NULL,
  total_price float(10) NOT NULL,
  delivery_time delivery_times NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(ID)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE orders;
DROP TYPE order_status;
DROP TYPE delivery_times;