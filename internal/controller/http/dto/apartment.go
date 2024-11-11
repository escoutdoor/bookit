package dto

import (
	"github.com/escoutdoor/bookit/pkg/validator"
)

type CreateApartmentDTO struct {
	Name        string  `json:"name" validate:"min=1"`
	Description *string `json:"description,omitempty"`
	Beds        *int    `json:"beds,omitempty" validate:"omitempty,gt=0"`
	Bedrooms    *int    `json:"bedrooms,omitempty" validate:"omitempty,gt=0"`
	Bathrooms   *int    `json:"bathrooms,omitempty" validate:"omitempty,gt=0"`
	MaxGuests   int     `json:"max_guests" validate:"gte=1"`
	RentalPrice float64 `json:"rental_price" validate:"required,gt=0"`
	Latitude    float64 `json:"latitude" validate:"required,latitude"`
	Longitude   float64 `json:"longitude" validate:"required,longitude"`
	CategoryID  *string `json:"category_id,omitempty" validate:"omitempty,uuid"`
}

type UpdateApartmentDTO struct {
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=1"`
	Description *string  `json:"description,omitempty"`
	Beds        *int     `json:"beds,omitempty" validate:"omitempty,gt=0"`
	Bedrooms    *int     `json:"bedrooms,omitempty" validate:"omitempty,gt=0"`
	Bathrooms   *int     `json:"bathrooms,omitempty" validate:"omitempty,gt=0"`
	MaxGuests   *int     `json:"max_guests,omitempty" validate:"omitempty,gte=1"`
	RentalPrice *float64 `json:"rental_price,omitempty" validate:"omitempty,gt=0"`
	Latitude    *float64 `json:"latitude,omitempty" validate:"omitempty,latitude"`
	Longitude   *float64 `json:"longitude,omitempty" validate:"omitempty,longitude"`
	CategoryID  *string  `json:"category_id,omitempty" validate:"omitempty,uuid"`
}

func (d *CreateApartmentDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	return v.Errm
}

func (d *UpdateApartmentDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	return v.Errm
}

type ApartmentQueryDTO struct {
	Name           *string  `json:"name,omitempty" validate:"omitempty,min=1"`
	Description    *string  `json:"description,omitempty"`
	Beds           *int     `json:"beds,omitempty" validate:"omitempty,gt=0"`
	Bedrooms       *int     `json:"bedrooms,omitempty" validate:"omitempty,gt=0"`
	Bathrooms      *int     `json:"bathrooms,omitempty" validate:"omitempty,gt=0"`
	MaxGuests      *int     `json:"max_guests,omitempty" validate:"omitempty,gte=1"`
	CategoryID     *string  `json:"category_id,omitempty" validate:"omitempty,uuid"`
	MinRentalPrice *float64 `json:"min_rentral_price,omitempty" validate:"omitempty,gt=0"`
	MaxRentalPrice *float64 `json:"max_rentral_price,omitempty" validate:"omitempty,gt=0"`

	Limit  int
	SortBy string
}

var apartSortFields = map[string]string{
	"date_added_desc":   "created_at desc",
	"date_added_asc":    "created_at asc",
	"rental_price_desc": "rental_price desc",
	"rental_price_asc":  "rental_price asc",
}

func (d *ApartmentQueryDTO) Validate() map[string]string {
	v := validator.New()

	if len(d.SortBy) > 0 {
		_, ok := apartSortFields[d.SortBy]
		if !ok {
			v.AddErr("sort_by", "invalid sort_by value")
		}
	}

	if d.MaxRentalPrice != nil || d.MinRentalPrice != nil {
		if *d.MinRentalPrice < *d.MaxRentalPrice {
			v.AddErr("max_rentral_price", "max_rentral_price should be greater than the minimum one")
		}
	}

	v.Validate(d)

	return v.Errm
}

func (d ApartmentQueryDTO) GetSortSQL() string {
	if d.SortBy == "" {
		return "created_at desc"
	}

	srt, ok := apartSortFields[d.SortBy]
	if !ok {
		return "created_at desc"
	}
	return srt
}

func (d ApartmentQueryDTO) GetLimit() int {
	if d.Limit == 0 {
		return defaultLimit
	}
	return d.Limit
}
