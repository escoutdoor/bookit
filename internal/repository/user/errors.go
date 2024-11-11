package user

import "errors"

var (
	ErrNotFound         = errors.New("user not found")
	ErrNoFieldsToUpdate = errors.New("no fields to update for user")

	ErrEmailAlreadyExists       = errors.New("user with this email already exists")
	ErrPhoneNumberAlreadyExists = errors.New("user with this phone number already exists")
)
