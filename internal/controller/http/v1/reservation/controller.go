package reservation

import (
	"log/slog"

	"github.com/escoutdoor/bookit/internal/usecase"
)

type controller struct {
	uc  usecase.ReservationUseCase
	log *slog.Logger
}

func NewReservationController(uc usecase.ReservationUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
