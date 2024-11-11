package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	"github.com/escoutdoor/bookit/pkg/hasher"
	"github.com/google/uuid"
)

func (uc *usecase) Update(ctx context.Context, in *model.UpdateUser, userID uuid.UUID) (*model.User, error) {
	const op = "UserUseCase.Update"

	if in.Password != nil {
		hashpw, err := hasher.HashPw(*in.Password)
		if err != nil {
			return nil, fmt.Errorf("%s: hash password error: %w", op, err)
		}
		in.Password = &hashpw
	}

	uu, err := uc.repo.Update(ctx, in, userID)
	if err != nil {
		switch {
		case errors.Is(err, userrepo.ErrEmailAlreadyExists):
			return nil, authuc.ErrEmailAlreadyExists
		case errors.Is(err, userrepo.ErrPhoneNumberAlreadyExists):
			return nil, authuc.ErrPhoneNumberAlreadyExists
		case errors.Is(err, userrepo.ErrNoFieldsToUpdate):
			return nil, ErrNoFieldsToUpdate
		default:
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}
	return uu, nil
}
