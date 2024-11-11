package apartment

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	repo  repository.ApartmentRepository
	cache repository.ApartmentCache
}

func NewApartmentUseCase(
	r repository.ApartmentRepository,
	cache repository.ApartmentCache,
) *usecase {
	return &usecase{
		repo:  r,
		cache: cache,
	}
}
