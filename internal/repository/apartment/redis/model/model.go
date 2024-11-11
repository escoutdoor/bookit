package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Apartment struct {
	ID          uuid.UUID  `redis:"id"`
	Name        string     `redis:"name"`
	Description string     `redis:"description"`
	Beds        *int       `redis:"beds"`
	Bedrooms    *int       `redis:"bedrooms"`
	Bathrooms   *int       `redis:"bathrooms"`
	MaxGuests   int        `redis:"max_guests"`
	RentalPrice float64    `redis:"rental_price"`
	Latitude    float64    `redis:"latitude"`
	Longitude   float64    `redis:"longitude"`
	HostID      uuid.UUID  `redis:"host_id"`
	CategoryID  *uuid.UUID `redis:"category_id"`
	CreatedAtNs int64      `redis:"created_at"`
}

func (a Apartment) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Apartment) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &a)
}
