package review

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	reviewRepo      repository.ReviewRepository
	reservationRepo repository.ReservationRepository
}

func NewReviewUseCase(
	reviewRepo repository.ReviewRepository,
	reservationRepo repository.ReservationRepository,
) *usecase {
	return &usecase{
		reviewRepo:      reviewRepo,
		reservationRepo: reservationRepo,
	}
}
