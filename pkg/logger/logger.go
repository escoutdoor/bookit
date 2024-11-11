package logger

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	prod  = "production"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case local:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true},
			),
		)
	case prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
