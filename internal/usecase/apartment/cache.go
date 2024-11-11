package apartment

import (
	"context"
	"errors"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	"github.com/google/uuid"
)

func (uc *usecase) getFromCache(ctx context.Context, id uuid.UUID) (*model.Apartment, error) {
	apart, err := uc.cache.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, aparterr.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	err = uc.cache.Expire(ctx, id)
	if err != nil {
		return nil, err
	}

	return apart, nil
}

func (uc *usecase) deleteFromCache(ctx context.Context, id uuid.UUID) error {
	err := uc.cache.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecase) setCache(ctx context.Context, v *model.Apartment) error {
	err := uc.cache.Set(ctx, v)
	if err != nil {
		return err
	}

	return nil
}
