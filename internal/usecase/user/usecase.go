package user

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *usecase {
	return &usecase{
		repo: r,
	}
}
