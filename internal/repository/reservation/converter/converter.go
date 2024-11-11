package converter

import (
	"github.com/escoutdoor/bookit/internal/model"
	repomodel "github.com/escoutdoor/bookit/internal/repository/reservation/model"
)

func ToReservationFromRepository(rsrv *repomodel.Reservation) *model.Reservation {
	return &model.Reservation{
		ID:          rsrv.ID,
		RenterID:    rsrv.RenterID,
		ApartmentID: rsrv.ApartmentID,
		StartDate:   rsrv.StartDate,
		EndDate:     rsrv.EndDate,
		Guests:      rsrv.Guests,
		Total:       rsrv.Total,
		CreatedAt:   rsrv.CreatedAt,
	}
}

func ToReservationsFromRepository(reporsrvs []*repomodel.Reservation) []*model.Reservation {
	var rsrvs []*model.Reservation
	for _, v := range reporsrvs {
		rsrvs = append(rsrvs, ToReservationFromRepository(v))
	}
	return rsrvs
}
