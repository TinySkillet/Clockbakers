-- +goose Up
CREATE TABLE password_reset_codes (
    id UUID PRIMARY KEY,
    email VARCHAR(64) NOT NULL,
    code VARCHAR(6) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_password_reset_codes_code ON password_reset_codes(code);
CREATE INDEX idx_password_reset_codes_email ON password_reset_codes(email);


-- +goose Down
DROP INDEX idx_password_reset_codes_code;
DROP INDEX idx_password_reset_codes_email;
DROP TABLE password_reset_codes;