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
	Rating    int32     `json:"rating" validate:"required,min=1,max=5"`
	Comment   string    `json:"comment" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
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

func DBReviewsToReviews(dbReviews []database.Review) []Review {
	reviews := make([]Review, len(dbReviews))
	for i, dbReview := range dbReviews {
		reviews[i] = DBReviewToReview(dbReview)
	}
	return reviews
}
