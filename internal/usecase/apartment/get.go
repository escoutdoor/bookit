package apartment

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	"github.com/google/uuid"
)

func (uc *usecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentUseCase.GetByID"

	apartment, err := uc.getFromCache(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: get apartment from cache: %s", op, err)
	}
	if apartment != nil {
		return apartment, nil
	}

	apartment, err = uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, aparterr.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	err = uc.setCache(ctx, apartment)
	if err != nil {
		return nil, fmt.Errorf("%s: set cache: %s", op, err)
	}

	return apartment, nil
}

func (uc *usecase) GetAll(ctx context.Context, in *model.ApartmentQuery) ([]*model.Apartment, error) {
	const op = "ApartmentUseCase.GetAll"
	aparts, err := uc.repo.GetAll(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return aparts, nil
}
