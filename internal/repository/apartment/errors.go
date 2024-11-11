package apartment

import "errors"

var (
	ErrNotFound         = errors.New("apartment not found")
	ErrNoFieldsToUpdate = errors.New("no fields to update for apartment")
)
