package review

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	"github.com/google/uuid"
)

func (uc *usecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Review, error) {
	const op = "ReviewUseCase.GetByID"
	rev, err := uc.reviewRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, categoryrepo.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return rev, nil
}

func (uc *usecase) GetAll(ctx context.Context, in *model.ReviewQuery) ([]*model.Review, error) {
	const op = "ReviewUseCase.GetAll"
	revs, err := uc.reviewRepo.GetAll(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return revs, nil
}
