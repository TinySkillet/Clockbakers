-- +goose Up
CREATE TABLE carts (
 ID UUID PRIMARY KEY NOT NULL,
 user_id UUID NOT NULL REFERENCES users(ID)
 ON DELETE CASCADE
);



-- +goose Down
DROP TABLE carts;