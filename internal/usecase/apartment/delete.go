package apartment

import (
	"context"
	"errors"
	"fmt"

	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	"github.com/google/uuid"
)

func (uc *usecase) Delete(ctx context.Context, apartmentID uuid.UUID, hostID uuid.UUID) error {
	const op = "ApartmentUseCase.Delete"
	apart, err := uc.repo.GetByID(ctx, apartmentID)
	if err != nil {
		if errors.Is(err, aparterr.ErrNotFound) {
			return ErrNotFound
		}

		return fmt.Errorf("%s: %s", op, err)
	}
	if apart.HostID != hostID {
		return authuc.ErrForbidden
	}

	err = uc.repo.Delete(ctx, apartmentID)
	if err != nil {
		return fmt.Errorf("%s: delete: %s", op, err)
	}

	err = uc.deleteFromCache(ctx, apartmentID)
	if err != nil {
		return fmt.Errorf("%s: delete from cache: %s", op, err)
	}

	return nil
}
