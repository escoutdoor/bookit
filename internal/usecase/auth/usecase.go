package auth

import (
	"time"

	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	repo     repository.UserRepository
	signKey  string
	tokenTTL time.Duration
}

func NewAuthUseCase(r repository.UserRepository, sk string, ttl time.Duration) *usecase {
	return &usecase{
		repo:     r,
		signKey:  sk,
		tokenTTL: ttl,
	}
}
