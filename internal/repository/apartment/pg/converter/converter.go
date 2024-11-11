package converter

import (
	"github.com/escoutdoor/bookit/internal/model"
	repomodel "github.com/escoutdoor/bookit/internal/repository/apartment/pg/model"
)

func ToApartmentFromRepository(apart *repomodel.Apartment) *model.Apartment {
	return &model.Apartment{
		ID:          apart.ID,
		Name:        apart.Name,
		Description: apart.Description,
		Beds:        apart.Beds,
		Bedrooms:    apart.Bedrooms,
		Bathrooms:   apart.Bathrooms,
		MaxGuests:   apart.MaxGuests,
		RentalPrice: apart.RentalPrice,
		Latitude:    apart.Latitude,
		Longitude:   apart.Longitude,
		HostID:      apart.HostID,
		CategoryID:  apart.CategoryID,
		CreatedAt:   apart.CreatedAt,
	}
}

func ToApartmentsFromRepository(repoaparts []*repomodel.Apartment) []*model.Apartment {
	var aparts []*model.Apartment
	for _, v := range repoaparts {
		aparts = append(aparts, ToApartmentFromRepository(v))
	}
	return aparts
}
