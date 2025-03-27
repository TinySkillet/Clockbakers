-- +goose Up
CREATE TABLE categories (
  name VARCHAR(32) PRIMARY KEY NOT NULL
);
INSERT INTO categories(
name
) VALUES (
'cake'
);

-- +goose Down
DROP TABLE categories;
