package avatar

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	repo repository.AvatarRepository
}

func NewAvatarUseCase(repo repository.AvatarRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
