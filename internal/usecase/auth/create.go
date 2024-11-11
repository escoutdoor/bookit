package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	"github.com/escoutdoor/bookit/pkg/hasher"
)

func (uc *usecase) Register(ctx context.Context, in *model.RegisterUser) (*model.AuthResponse, error) {
	const op = "AuthUseCase.Register"
	var err error

	in.Password, err = hasher.HashPw(in.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: hash password error: %w", op, err)
	}

	u, err := uc.repo.Create(ctx, in)
	if err != nil {
		switch {
		case errors.Is(err, userrepo.ErrEmailAlreadyExists):
			return nil, ErrEmailAlreadyExists
		case errors.Is(err, userrepo.ErrPhoneNumberAlreadyExists):
			return nil, ErrPhoneNumberAlreadyExists
		default:
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}

	token, err := uc.genToken(u)
	if err != nil {
		return nil, fmt.Errorf("%s: generate token error: %w", op, err)
	}

	resp := &model.AuthResponse{
		User:        u,
		AccessToken: token,
	}
	return resp, nil
}

func (uc *usecase) Login(ctx context.Context, in *model.Login) (*model.AuthResponse, error) {
	const op = "AuthUseCase.Login"

	u, err := uc.repo.GetByEmail(ctx, in.Email)
	if err != nil {
		if errors.Is(err, userrepo.ErrNotFound) {
			return nil, ErrInvalidEmailOrPassword
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	eq := hasher.ComparePw(in.Password, u.Password)
	if !eq {
		return nil, ErrInvalidEmailOrPassword
	}

	token, err := uc.genToken(u)
	if err != nil {
		return nil, fmt.Errorf("%s: generate token error: %w", op, err)
	}

	resp := &model.AuthResponse{
		User:        u,
		AccessToken: token,
	}
	return resp, nil
}
