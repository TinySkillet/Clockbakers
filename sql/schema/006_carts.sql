-- +goose Up
CREATE TABLE carts (
 ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
 user_id UUID NOT NULL REFERENCES users(ID)
 ON DELETE CASCADE
);

INSERT INTO carts (
  ID, user_id )
VALUES ('g30ac10e-58cc-4392-a587-0e03b2c3d480', 'f47ac10b-58cc-4372-a567-0e02b2c3d479');

-- +goose Down
DROP TABLE carts;