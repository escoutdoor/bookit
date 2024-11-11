package avatar

import (
	"log/slog"

	"github.com/escoutdoor/bookit/internal/usecase"
)

type controller struct {
	uc  usecase.AvatarUseCase
	log *slog.Logger
}

func NewAvatarController(uc usecase.AvatarUseCase, log *slog.Logger) *controller {
	return &controller{
		uc:  uc,
		log: log,
	}
}
