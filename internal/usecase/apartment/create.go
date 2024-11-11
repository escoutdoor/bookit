package apartment

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
	useruc "github.com/escoutdoor/bookit/internal/usecase/user"
	"github.com/google/uuid"
)

func (uc *usecase) Create(ctx context.Context, in *model.CreateApartment, hostID uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentUseCase.Create"
	apart, err := uc.repo.Create(ctx, in, hostID)
	if err != nil {
		switch {
		case errors.Is(err, categoryrepo.ErrNotFound):
			return nil, categoryuc.ErrNotFound
		case errors.Is(err, userrepo.ErrNotFound):
			return nil, useruc.ErrNotFound
		default:
			return nil, fmt.Errorf("%s: %s", op, err)
		}
	}

	return apart, nil
}
