-- name: CreatePasswordResetCode :one
INSERT INTO password_reset_codes (id, email, code, expires_at, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetValidPasswordResetCode :one
SELECT * FROM password_reset_codes
WHERE email = $1 AND expires_at > CURRENT_TIMESTAMP
ORDER BY created_at DESC
LIMIT 1;

-- name: DeleteExpiredPasswordResetCodes :exec
DELETE FROM password_reset_codes
WHERE expires_at <= CURRENT_TIMESTAMP;