package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/escoutdoor/bookit/internal/config"
)

type Server interface {
	Run() error
	GracefulStop()
}

type server struct {
	httpSrv *http.Server
}

func New(handler http.Handler, cfg *config.HTTPSrvConfig) Server {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        handler,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return &server{
		httpSrv: srv,
	}
}

func (s *server) Run() error {
	if err := s.httpSrv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *server) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	s.httpSrv.Shutdown(ctx)
}
