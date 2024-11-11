package category

import "errors"

var (
	ErrNotFound         = errors.New("category not found")
	ErrNoFieldsToUpdate = errors.New("no fields to update for category")
)
