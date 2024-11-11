package user

import "errors"

var (
	ErrNotFound         = errors.New("user not found")
	ErrNoFieldsToUpdate = errors.New("no fields to update for user")
)
