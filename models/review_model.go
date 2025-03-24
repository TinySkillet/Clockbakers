package models

import (
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type GetReviewsParams struct {
	ID        uuid.NullUUID `json:"id"`
	ProductID uuid.NullUUID `json:"product_id"`
	UserId    uuid.NullUUID `json:"user_id"`
}

type Review struct {
	ID        uuid.UUID `json:"id"`
	Rating    int32     `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
}

func DBReviewToReview(dbReview database.Review) Review {
	return Review{
		ID:        dbReview.ID,
		Rating:    dbReview.Rating,
		Comment:   dbReview.Comment,
		CreatedAt: dbReview.CreatedAt,
		UserID:    dbReview.UserID,
		ProductID: dbReview.ProductID,
	}
}
