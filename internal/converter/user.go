package converter

import (
	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/pkg/validator"
)

func ToUpdateUserFromDTO(dto *dto.UpdateUserDTO) (*model.UpdateUser, error) {
	v := validator.New()
	uu := &model.UpdateUser{
		Email:       dto.Email,
		Password:    dto.Password,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		AvatarURL:   dto.AvatarURL,
		PhoneNumber: dto.PhoneNumber,
	}

	if dto.DOB != nil {
		dob, err := v.ParseDate(*dto.DOB)
		if err != nil {
			return nil, err
		}
		uu.DOB = &dob
	}

	return uu, nil
}
