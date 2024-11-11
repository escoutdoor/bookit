package auth

import (
	"errors"
)

var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")

	ErrEmailAlreadyExists       = errors.New("user with this email already exists")
	ErrPhoneNumberAlreadyExists = errors.New("user with this phone number already exists")

	ErrForbidden = errors.New("forbidden")
)
