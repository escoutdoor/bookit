package model

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID          uuid.UUID `json:"id"`
	RenterID    uuid.UUID `json:"renter_id"`
	ApartmentID uuid.UUID `json:"apartment_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Guests      int       `json:"guests"`
	Total       float64   `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateReservation struct {
	ApartmentID uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
	Guests      int
}

type UpdateReservation struct {
	StartDate *time.Time
	EndDate   *time.Time
	Guests    *int
	Total     *float64
}

type ReservationQuery struct {
	ApartmentID *uuid.UUID
	StartDate   *time.Time
	EndDate     *time.Time

	Limit int
}
