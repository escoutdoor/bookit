package reservation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	"github.com/google/uuid"
)

func (uc *usecase) Create(ctx context.Context, in *model.CreateReservation, renterID uuid.UUID) (*model.Reservation, error) {
	const op = "ReservationUseCase.Create"

	apart, err := uc.apartRepo.GetByID(ctx, in.ApartmentID)
	if err != nil {
		if errors.Is(err, aparterr.ErrNotFound) {
			return nil, apartmentuc.ErrNotFound
		}

		return nil, fmt.Errorf("%s: %s", op, err)
	}

	total, err := uc.countTotal(apart.RentalPrice, in.StartDate, in.EndDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	ib, err := uc.checkOverlap(ctx, in.ApartmentID, in.StartDate, in.EndDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	if ib {
		return nil, ErrAlreadyExists
	}

	rsrv, err := uc.rsrvRepo.Create(ctx, in, renterID, total)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return rsrv, nil
}

func (uc *usecase) countTotal(price float64, start, end time.Time) (float64, error) {
	diff := int(end.Sub(start).Hours() / 24)
	return price * float64(diff), nil
}

func (uc *usecase) checkOverlap(ctx context.Context, apartmentID uuid.UUID, startDate, endDate time.Time) (bool, error) {
	const op = "ReservationUseCase.checkOverlap"
	rsrvs, err := uc.rsrvRepo.GetBetweenDates(ctx, apartmentID, startDate, endDate)
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	}

	return len(rsrvs) > 0, nil
}
