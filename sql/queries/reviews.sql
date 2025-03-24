-- name: CreateReview :one
INSERT INTO reviews (id, rating, comment, created_at, updated_at, user_id, product_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetReviews :many
SELECT * FROM reviews
WHERE 
  ($1::UUID IS NULL OR id = $1) AND
  ($2::UUID IS NULL OR product_id = $2) AND
  ($3::UUID IS NULL OR user_id = $3)
ORDER BY created_at DESC;

-- name: UpdateReview :one
UPDATE reviews
SET rating = $3, comment = $4, updated_at = $5
WHERE id = $1
AND user_id = $2
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews WHERE id = $1 AND user_id = $2;