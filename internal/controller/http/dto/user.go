package dto

import (
	"github.com/escoutdoor/bookit/pkg/validator"
)

type UpdateUserDTO struct {
	Email       *string `json:"email,omitempty" validate:"omitempty,email"`
	Password    *string `json:"password,omitempty" validate:"omitempty,min=6,max=30,password"`
	FirstName   *string `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName    *string `json:"last_name,omitempty" validate:"omitempty,min=2"`
	DOB         *string `json:"date_of_birth,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty" validate:"omitempty,url"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,e164"`
}

func (d *UpdateUserDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	if d.DOB != nil {
		_, err := v.ParseDate(*d.DOB)
		if err != nil {
			v.AddErr("date_of_birth", err.Error())
		}
	}

	return v.Errm
}
