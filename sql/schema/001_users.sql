-- +goose Up
CREATE TYPE user_type AS ENUM ('customer', 'admin');
CREATE TABLE users(
  ID UUID PRIMARY KEY NOT NULL,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL UNIQUE,
  phone_no VARCHAR(64),
  address TEXT,
  password VARCHAR(64) NOT NULL,
  role user_type NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
  );

INSERT INTO users (
    ID, first_name, last_name,
    email, phone_no,
    address, password, role,
    created_at, updated_at
) VALUES (
    'f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Subarna', 'Bajracharya',
    'subarna@example.com', '+1234567890',
    'Gabahal, Lalitpur',
    '3e94d2611ea3c3b08bb705820850fc9a4725fba58553101cdf256f0b4a660c2d', 'admin',
    CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
);
  
-- +goose Down
DROP TABLE users;
DROP TYPE user_type;
