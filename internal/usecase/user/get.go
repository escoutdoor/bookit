package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	"github.com/google/uuid"
)

func (uc *usecase) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	const op = "UserUseCase.GetByID"

	u, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, userrepo.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return u, nil
}
