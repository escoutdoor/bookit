package auth

import (
	"github.com/escoutdoor/bookit/internal/usecase"
	"log/slog"
)

type controller struct {
	uc  usecase.AuthUseCase
	log *slog.Logger
}

func NewAuthController(uc usecase.AuthUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
