package reservation

import "errors"

var (
	ErrNotFound         = errors.New("reservation not found")
	ErrNoFieldsToUpdate = errors.New("no fields to update for reservation")
)
