package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Role        string    `json:"role"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DOB         time.Time `json:"date_of_birth"`
	AvatarURL   *string   `json:"avatar_url"`
	PhoneNumber *string   `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateUser struct {
	Email       *string
	Password    *string
	FirstName   *string
	LastName    *string
	DOB         *time.Time
	AvatarURL   *string
	PhoneNumber *string
}
