package category

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	"github.com/google/uuid"
)

func (uc *usecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	const op = "CategoryUseCase.GetByID"

	c, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, categoryrepo.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return c, nil
}
