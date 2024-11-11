package model

import (
	"time"

	"github.com/google/uuid"
)

type Apartment struct {
	ID          uuid.UUID
	Name        string
	Description string
	Beds        *int
	Bedrooms    *int
	Bathrooms   *int
	MaxGuests   int
	RentalPrice float64
	Latitude    float64
	Longitude   float64
	HostID      uuid.UUID
	CategoryID  *uuid.UUID
	CreatedAt   time.Time
}
