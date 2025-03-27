-- +goose Up
CREATE TABLE delivery_addresses (
  ID UUID NOT NULL PRIMARY KEY,
  address TEXT NOT NULL,
  user_id UUID REFERENCES users(ID) NOT NULL
);

-- +goose Down
DROP TABLE delivery_addresses;