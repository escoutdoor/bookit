package review

import (
	"log/slog"

	"github.com/escoutdoor/bookit/internal/usecase"
)

type controller struct {
	uc  usecase.ReviewUseCase
	log *slog.Logger
}

func NewReviewController(uc usecase.ReviewUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
