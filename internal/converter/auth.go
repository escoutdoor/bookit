package converter

import (
	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
)

func ToRegisterUserFromDTO(dto *dto.RegisterUserDTO) *model.RegisterUser {
	return &model.RegisterUser{
		Email:     dto.Email,
		Password:  dto.Password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		DOB:       dto.DOB,
	}
}

func ToLoginFromDTO(dto *dto.LoginDTO) *model.Login {
	return &model.Login{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
