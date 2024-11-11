package apartment

import (
	"log/slog"

	"github.com/escoutdoor/bookit/internal/usecase"
)

type controller struct {
	uc  usecase.ApartmentUseCase
	log *slog.Logger
}

func NewApartmentController(uc usecase.ApartmentUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
