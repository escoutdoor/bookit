package review

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	useruc "github.com/escoutdoor/bookit/internal/usecase/user"
	"github.com/google/uuid"
)

func (uc *usecase) Create(ctx context.Context, in *model.CreateReview, renterID uuid.UUID) (*model.Review, error) {
	const op = "ReviewUseCase.Create"

	isRentedBefore, err := uc.reservationRepo.IsRentedBefore(ctx, in.ApartmentID, renterID)
	if err != nil {
		switch {
		case errors.Is(err, aparterr.ErrNotFound):
			return nil, apartmentuc.ErrNotFound
		case errors.Is(err, userrepo.ErrNotFound):
			return nil, useruc.ErrNotFound
		}
		return nil, fmt.Errorf("%s: check is rented before: %s", op, err)
	}
	if !isRentedBefore {
		return nil, ErrNotAllowedToCreateReview
	}

	rev, err := uc.reviewRepo.Create(ctx, in, renterID)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return rev, nil
}
