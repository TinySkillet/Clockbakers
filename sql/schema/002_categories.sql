-- +goose Up
CREATE TABLE categories (
  name VARCHAR(32) PRIMARY KEY NOT NULL
);


-- +goose Down
DROP TABLE categories;
