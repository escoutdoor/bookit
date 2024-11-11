package reservation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	reservationrepo "github.com/escoutdoor/bookit/internal/repository/reservation"
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	"github.com/google/uuid"
)

func (uc *usecase) Update(ctx context.Context, in *model.UpdateReservation, rsrvID, renterID uuid.UUID) (*model.Reservation, error) {
	const op = "ReservationUseCase.Update"
	rsrv, err := uc.GetByID(ctx, rsrvID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s: get reservation by id error %s", op, err)
	}
	if rsrv.RenterID != renterID {
		return nil, authuc.ErrForbidden
	}

	if in.StartDate != nil || in.EndDate != nil {
		if in.StartDate != nil {
			rsrv.StartDate = *in.StartDate
		}
		if in.EndDate != nil {
			rsrv.EndDate = *in.EndDate
		}
		if err = uc.validateRsrvRange(rsrv.StartDate, rsrv.EndDate); err != nil {
			return nil, err
		}

		apart, err := uc.apartRepo.GetByID(ctx, rsrv.ApartmentID)
		if err != nil {
			if errors.Is(err, aparterr.ErrNotFound) {
				return nil, apartmentuc.ErrNotFound
			}
			return nil, fmt.Errorf("%s: %s", op, err)
		}

		ib, err := uc.checkUpdateOverlap(ctx, rsrvID, apart.ID, rsrv.StartDate, rsrv.EndDate)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		if ib {
			return nil, ErrAlreadyExists
		}

		total, err := uc.countTotal(apart.RentalPrice, rsrv.StartDate, rsrv.EndDate)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		if total > 0 {
			in.Total = &total
		}
	}

	ursrv, err := uc.rsrvRepo.Update(ctx, in, rsrvID)
	if err != nil {
		if errors.Is(err, reservationrepo.ErrNoFieldsToUpdate) {
			return nil, ErrNoFieldsToUpdate
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return ursrv, nil
}

func (uc *usecase) validateRsrvRange(startDate, endDate time.Time) error {
	if startDate.After(endDate) {
		return ErrInvalidReservationRange
	}
	return nil
}

func (uc *usecase) checkUpdateOverlap(ctx context.Context, rsrvID, apartmentID uuid.UUID, startDate, endDate time.Time) (bool, error) {
	const op = "ReservationUseCase.checkUpdateOverlap"
	rsrvs, err := uc.rsrvRepo.GetBetweenDatesExcludingID(ctx, rsrvID, apartmentID, startDate, endDate)
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	}

	return len(rsrvs) > 0, nil
}
