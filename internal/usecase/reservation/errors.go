package reservation

import (
	"errors"
)

var (
	ErrNotFound                = errors.New("reservation not found")
	ErrAlreadyExists           = errors.New("reservation already exists")
	ErrNoFieldsToUpdate        = errors.New("no fields to update for reservation")
	ErrInvalidReservationRange = errors.New("the end date must be after the start date")
)
