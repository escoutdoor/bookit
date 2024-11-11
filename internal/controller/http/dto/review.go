package dto

import (
	"github.com/escoutdoor/bookit/pkg/validator"
)

type CreateReviewDTO struct {
	Content     string `json:"content" validate:"min=5"`
	Rating      int    `json:"rating" validate:"gt=0,lte=5"`
	ApartmentID string `json:"apartment_id" validate:"required,uuid"`
}

type ReviewQueryDTO struct {
	Rating      *int    `json:"rating,omitempty" validate:"omitempty,gt=0,lte=5"`
	RenterID    *string `json:"renter_id,omitempty" validate:"omitempty,uuid"`
	ApartmentID *string `json:"apartment_id,omitempty" validate:"omitempty,uuid"`

	Limit int
}

func (d *CreateReviewDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)
	return v.Errm
}

func (d ReviewQueryDTO) GetLimit() int {
	if d.Limit == 0 {
		return defaultLimit
	}
	return d.Limit
}
