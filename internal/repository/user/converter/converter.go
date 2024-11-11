package converter

import (
	"github.com/escoutdoor/bookit/internal/model"
	repomodel "github.com/escoutdoor/bookit/internal/repository/user/model"
)

func ToUserFromRepository(u *repomodel.User) *model.User {
	return &model.User{
		ID:          u.ID,
		Email:       u.Email,
		Password:    u.Password,
		Role:        u.Role,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		DOB:         u.DOB,
		AvatarURL:   u.AvatarURL,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
	}
}
