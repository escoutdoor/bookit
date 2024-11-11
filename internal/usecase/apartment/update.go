package apartment

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
	"github.com/google/uuid"
)

func (uc *usecase) Update(ctx context.Context, in *model.UpdateApartment, apartmentID uuid.UUID, hostID uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentUseCase.Update"

	apart, err := uc.repo.GetByID(ctx, apartmentID)
	if err != nil {
		if errors.Is(err, aparterr.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}
	if apart.HostID != hostID {
		return nil, authuc.ErrForbidden
	}

	uapart, err := uc.repo.Update(ctx, in, apartmentID)
	if err != nil {
		switch {
		case errors.Is(err, categoryrepo.ErrNotFound):
			return nil, categoryuc.ErrNotFound
		case errors.Is(err, aparterr.ErrNoFieldsToUpdate):
			return nil, ErrNoFieldsToUpdate
		default:
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}

	fmt.Printf("[CACHE] REMOVED FROM CACHE ID: %s\n", uapart.ID)
	err = uc.cache.Delete(ctx, apartmentID)
	if err != nil {
		return nil, fmt.Errorf("%s: delete from cache: %s", op, err)
	}

	return uapart, nil
}
