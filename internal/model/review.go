package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `json:"id"`
	Content     string    `json:"content"`
	Rating      int       `json:"rating"`
	RenterID    uuid.UUID `json:"renter_id"`
	ApartmentID uuid.UUID `json:"apartment_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateReview struct {
	Content     string
	Rating      int
	ApartmentID uuid.UUID
}

type ReviewQuery struct {
	Rating      *int
	RenterID    *uuid.UUID
	ApartmentID *uuid.UUID

	Limit int
}
