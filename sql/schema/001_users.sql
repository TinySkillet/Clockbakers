-- +goose Up
CREATE TYPE user_type AS ENUM ('customer', 'admin');
CREATE TABLE users(
  ID UUID PRIMARY KEY NOT NULL,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL UNIQUE,
  phone_no VARCHAR(64) NOT NULL,
  address TEXT NOT NULL,
  password VARCHAR(64) NOT NULL,
  role user_type NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
  );
  
-- +goose Down
DROP TABLE users;
DROP TYPE user_type;
