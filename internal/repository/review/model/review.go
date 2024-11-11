package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID
	Content     string
	Rating      int
	RenterID    uuid.UUID
	ApartmentID uuid.UUID
	CreatedAt   time.Time
}
