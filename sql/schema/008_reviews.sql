-- +goose Up
CREATE TABLE reviews (
  ID UUID PRIMARY KEY NOT NULL,
  rating INTEGER NOT NULL,
  comment TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(ID),
  product_id UUID NOT NULL REFERENCES products(ID) 
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE reviews;