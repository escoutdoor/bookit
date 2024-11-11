package model

import (
	"time"

	"github.com/google/uuid"
)

type Apartment struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Beds        *int       `json:"beds"`
	Bedrooms    *int       `json:"bedrooms"`
	Bathrooms   *int       `json:"bathrooms"`
	MaxGuests   int        `json:"max_guests"`
	RentalPrice float64    `json:"rental_price"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	HostID      uuid.UUID  `json:"host_id"`
	CategoryID  *uuid.UUID `json:"category_id"`

	CreatedAt time.Time `json:"created_at"`
}

type CreateApartment struct {
	Name        string
	Description *string
	Beds        *int
	Bedrooms    *int
	Bathrooms   *int
	MaxGuests   int
	RentalPrice float64
	Latitude    float64
	Longitude   float64
	CategoryID  *uuid.UUID
}

type UpdateApartment struct {
	Name        *string
	Description *string
	Beds        *int
	Bedrooms    *int
	Bathrooms   *int
	MaxGuests   *int
	RentalPrice *float64
	Latitude    *float64
	Longitude   *float64
	CategoryID  *uuid.UUID
}

type ApartmentQuery struct {
	Name           *string
	Description    *string
	Beds           *int
	Bedrooms       *int
	Bathrooms      *int
	MaxGuests      *int
	CategoryID     *uuid.UUID
	MinRentalPrice *float64
	MaxRentalPrice *float64

	Limit  int
	SortBy string
}
