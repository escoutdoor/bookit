package category

import (
	"context"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
)

func (uc *usecase) Create(ctx context.Context, in *model.CreateCategory) (*model.Category, error) {
	const op = "CategoryUseCase.Create"

	category, err := uc.repo.Create(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return category, err
}
