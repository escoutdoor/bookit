package model

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID          uuid.UUID
	RenterID    uuid.UUID
	ApartmentID uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
	Guests      int
	Total       float64
	CreatedAt   time.Time
}
