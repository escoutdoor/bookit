package converter

import (
	"fmt"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
)

func ToCreateApartmentFromDTO(dto *dto.CreateApartmentDTO) (*model.CreateApartment, error) {
	ca := &model.CreateApartment{
		Name:        dto.Name,
		Description: dto.Description,
		Beds:        dto.Beds,
		Bedrooms:    dto.Bedrooms,
		Bathrooms:   dto.Bathrooms,
		MaxGuests:   dto.MaxGuests,
		RentalPrice: dto.RentalPrice,
		Latitude:    dto.Latitude,
		Longitude:   dto.Longitude,
	}
	if dto.CategoryID != nil {
		ctID, err := uuid.Parse(*dto.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("invalid category id, must be UUID")
		}
		ca.CategoryID = &ctID
	}

	return ca, nil
}

func ToUpdateApartmentFromDTO(dto *dto.UpdateApartmentDTO) (*model.UpdateApartment, error) {
	mua := &model.UpdateApartment{
		Name:        dto.Name,
		Description: dto.Description,
		Beds:        dto.Beds,
		Bedrooms:    dto.Bedrooms,
		Bathrooms:   dto.Bathrooms,
		MaxGuests:   dto.MaxGuests,
		RentalPrice: dto.RentalPrice,
		Latitude:    dto.Latitude,
		Longitude:   dto.Longitude,
	}
	if dto.CategoryID != nil {
		ctID, err := uuid.Parse(*dto.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("invalid category id, must be UUID")
		}
		mua.CategoryID = &ctID
	}

	return mua, nil
}

func ToApartmentQueryFromDTO(dto *dto.ApartmentQueryDTO) (*model.ApartmentQuery, error) {
	aq := &model.ApartmentQuery{
		Name:           dto.Name,
		Description:    dto.Description,
		Beds:           dto.Beds,
		Bedrooms:       dto.Bedrooms,
		Bathrooms:      dto.Bathrooms,
		MaxGuests:      dto.MaxGuests,
		MinRentalPrice: dto.MinRentalPrice,
		MaxRentalPrice: dto.MaxRentalPrice,

		SortBy: dto.GetSortSQL(),
		Limit:  dto.GetLimit(),
	}

	if dto.CategoryID != nil {
		ctID, err := uuid.Parse(*dto.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("invalid category id, must be UUID")
		}
		aq.CategoryID = &ctID
	}

	return aq, nil
}
