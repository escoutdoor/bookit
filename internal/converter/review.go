package converter

import (
	"fmt"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
)

func ToCreateReviewFromDTO(dto *dto.CreateReviewDTO) (*model.CreateReview, error) {
	apartmentID, err := uuid.Parse(dto.ApartmentID)
	if err != nil {
		return nil, fmt.Errorf("invalid apartment id, must be UUID")
	}

	return &model.CreateReview{
		Content:     dto.Content,
		Rating:      dto.Rating,
		ApartmentID: apartmentID,
	}, nil
}

func ToReviewQueryFromDTO(dto *dto.ReviewQueryDTO) (*model.ReviewQuery, error) {
	rq := &model.ReviewQuery{
		Limit: dto.GetLimit(),
	}

	if dto.ApartmentID != nil {
		apartmentID, err := uuid.Parse(*dto.ApartmentID)
		if err != nil {
			return nil, fmt.Errorf("invalid apartment id, must be UUID")
		}
		rq.ApartmentID = &apartmentID
	}

	if dto.RenterID != nil {
		renterID, err := uuid.Parse(*dto.RenterID)
		if err != nil {
			return nil, fmt.Errorf("invalid renter id, must be UUID")
		}
		rq.RenterID = &renterID
	}

	if dto.Rating != nil {
		if *dto.Rating <= 0 || *dto.Rating > 5 {
			return nil, fmt.Errorf("rating is out of range, expected from 1 to 5")
		}
		rq.Rating = dto.Rating
	}

	return rq, nil
}
