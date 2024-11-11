package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       string
	Password    string
	Role        string
	FirstName   string
	LastName    string
	DOB         time.Time
	AvatarURL   *string
	PhoneNumber *string
	CreatedAt   time.Time
}
