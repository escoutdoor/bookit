package reservation

import (
	"github.com/escoutdoor/bookit/internal/repository"
)

type usecase struct {
	rsrvRepo  repository.ReservationRepository
	apartRepo repository.ApartmentRepository
}

func NewReservationUseCase(
	rsrvRepo repository.ReservationRepository,
	apartRepo repository.ApartmentRepository,
) *usecase {
	return &usecase{
		rsrvRepo:  rsrvRepo,
		apartRepo: apartRepo,
	}
}
