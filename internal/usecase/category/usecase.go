package category

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(r repository.CategoryRepository) *usecase {
	return &usecase{
		repo: r,
	}
}
