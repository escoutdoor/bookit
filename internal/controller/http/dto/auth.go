package dto

import (
	"github.com/escoutdoor/bookit/pkg/validator"
)

type RegisterUserDTO struct {
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min=6,max=30,password"`
	FirstName string `json:"first_name" validate:"min=2"`
	LastName  string `json:"last_name" validate:"min=2"`
	DOB       string `json:"date_of_birth" validate:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6,max=30,password"`
}

func (d *RegisterUserDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	_, err := v.ParseDate(d.DOB)
	if err != nil {
		v.AddErr("date_of_birth", err.Error())
	}
	return v.Errm
}

func (d *LoginDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	return v.Errm
}
