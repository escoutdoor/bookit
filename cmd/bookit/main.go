package main

import (
	"context"
	"os"

	"github.com/escoutdoor/bookit/internal/app"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/escoutdoor/bookit/pkg/logger"
)

const (
	cfgPath = "config/config.yaml"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load(cfgPath)
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.App.Env)

	err = app.Run(ctx, cfg, log)
	if err != nil {
		log.Error("run application", "error", err)
		os.Exit(1)
	}
}
