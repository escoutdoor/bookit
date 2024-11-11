package reservation

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	reservationrepo "github.com/escoutdoor/bookit/internal/repository/reservation"
	"github.com/google/uuid"
)

func (uc *usecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Reservation, error) {
	const op = "ReservationUseCase.GetByID"
	rsrv, err := uc.rsrvRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, reservationrepo.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return rsrv, nil
}

func (uc *usecase) GetAll(ctx context.Context, in *model.ReservationQuery) ([]*model.Reservation, error) {
	const op = "ReservationUseCase.GetAll"
	rsrvs, err := uc.rsrvRepo.GetAll(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return rsrvs, nil
}
