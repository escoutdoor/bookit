package category

import (
	"log/slog"

	"github.com/escoutdoor/bookit/internal/usecase"
)

type controller struct {
	uc  usecase.CategoryUseCase
	log *slog.Logger
}

func NewCategoryController(uc usecase.CategoryUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
