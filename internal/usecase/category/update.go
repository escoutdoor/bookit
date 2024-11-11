package category

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	"github.com/google/uuid"
)

func (uc *usecase) Update(ctx context.Context, in *model.UpdateCategory, categoryID uuid.UUID) (*model.Category, error) {
	const op = "CategoryUseCase.Update"

	_, err := uc.GetByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: get by id %s", op, err)
	}

	category, err := uc.repo.Update(ctx, in, categoryID)
	if err != nil {
		if errors.Is(err, categoryrepo.ErrNoFieldsToUpdate) {
			return nil, ErrNoFieldsToUpdate
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return category, nil
}
