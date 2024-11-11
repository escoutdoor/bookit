package dto

import (
	"time"

	"github.com/escoutdoor/bookit/pkg/validator"
)

type CreateReservationDTO struct {
	ApartmentID string `json:"apartment_id" validate:"uuid"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	Guests      int    `json:"guests" validate:"required,gte=1"`
}

func (d *CreateReservationDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	start, err := v.ParseDate(d.StartDate)
	if err != nil {
		v.AddErr("start_date", err.Error())
	}

	end, err := v.ParseDate(d.EndDate)
	if err != nil {
		v.AddErr("end_date", err.Error())
	}

	if start.After(end) {
		v.AddErr("reservation_range", "the end date must be after the start date")
	}
	return v.Errm
}

type UpdateReservationDTO struct {
	StartDate *string `json:"start_date,omitempty"`
	EndDate   *string `json:"end_date,omitempty"`
	Guests    *int    `json:"guests,omitempty" validate:"omitempty,gte=1"`
}

func (d *UpdateReservationDTO) Validate() map[string]string {
	var (
		v          = validator.New()
		start, end time.Time
		err        error
	)
	v.Validate(d)

	if d.StartDate != nil {
		start, err = v.ParseDate(*d.StartDate)
		if err != nil {
			v.AddErr("start_date", err.Error())
		}
	}

	if d.EndDate != nil {
		end, err = v.ParseDate(*d.EndDate)
		if err != nil {
			v.AddErr("end_date", err.Error())
		}
	}

	if d.EndDate != nil && d.StartDate != nil {
		if start.After(end) {
			v.AddErr("reservation_range", "the end date must be after the start date")
		}
	}
	return v.Errm
}

type ReservationQueryDTO struct {
	ApartmentID *string `json:"apartment_id,omitempty" validate:"omitempty,uuid"`
	StartDate   *string `json:"start_date,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`

	Limit int
}

func (d *ReservationQueryDTO) Validate() map[string]string {
	var (
		v          = validator.New()
		start, end time.Time
		err        error
	)
	v.Validate(d)

	if d.StartDate != nil {
		start, err = v.ParseDate(*d.StartDate)
		if err != nil {
			v.AddErr("start_date", err.Error())
		}
	}

	if d.EndDate != nil {
		end, err = v.ParseDate(*d.EndDate)
		if err != nil {
			v.AddErr("end_date", err.Error())
		}
	}

	if d.EndDate != nil && d.StartDate != nil {
		if start.After(end) {
			v.AddErr("reservation_range", "the end date must be after the start date")
		}
	}
	return v.Errm
}

func (d ReservationQueryDTO) GetLimit() int {
	if d.Limit == 0 {
		return defaultLimit
	}
	return d.Limit
}
