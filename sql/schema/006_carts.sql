-- +goose Up
CREATE TABLE carts (
 ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
 user_id UUID NOT NULL REFERENCES users(ID)
 ON DELETE CASCADE
);

INSERT INTO carts (
  ID, user_id )
VALUES ('7f3a1e4c-92f5-4e8b-9d2a-1b3f5e7d9c0b', 'f47ac10b-58cc-4372-a567-0e02b2c3d479');

-- +goose Down
DROP TABLE carts;