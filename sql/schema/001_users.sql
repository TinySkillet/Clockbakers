-- +goose Up
CREATE TABLE users(
  ID UUID PRIMARY KEY NOT NULL,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL UNIQUE,
  phone_no VARCHAR(64) NOT NULL,
  address TEXT NOT NULL,
  password VARCHAR(64) NOT NULL,
  role VARCHAR(16) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
  );
  
-- +goose Down
DROP TABLE users;