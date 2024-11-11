package dto

import (
	"github.com/escoutdoor/bookit/pkg/validator"
)

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=2"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name,omitempty" validate:"omitempty,min=2"`
}

func (d *CreateCategoryDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	return v.Errm
}

func (d *UpdateCategoryDTO) Validate() map[string]string {
	v := validator.New()
	v.Validate(d)

	return v.Errm
}
