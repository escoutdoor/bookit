package user

import (
	"github.com/escoutdoor/bookit/internal/usecase"
	"log/slog"
)

type controller struct {
	uc  usecase.UserUseCase
	log *slog.Logger
}

func NewUserController(uc usecase.UserUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
