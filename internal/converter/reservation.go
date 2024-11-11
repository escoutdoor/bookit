package converter

import (
	"fmt"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/pkg/validator"
	"github.com/google/uuid"
)

func ToCreateReservationFromDTO(dto *dto.CreateReservationDTO) (*model.CreateReservation, error) {
	apartmentID, err := uuid.Parse(dto.ApartmentID)
	if err != nil {
		return nil, fmt.Errorf("invalid apartment id, must be UUID")
	}

	v := validator.New()
	startDate, err := v.ParseDate(dto.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := v.ParseDate(dto.EndDate)
	if err != nil {
		return nil, err
	}

	return &model.CreateReservation{
		ApartmentID: apartmentID,
		StartDate:   startDate,
		EndDate:     endDate,
		Guests:      dto.Guests,
	}, nil
}

func ToUpdateReservationFromDTO(dto *dto.UpdateReservationDTO) (*model.UpdateReservation, error) {
	ur := &model.UpdateReservation{}
	v := validator.New()

	if dto.StartDate != nil {
		startDate, err := v.ParseDate(*dto.StartDate)
		if err != nil {
			return nil, err
		}
		ur.StartDate = &startDate
	}
	if dto.EndDate != nil {
		endDate, err := v.ParseDate(*dto.EndDate)
		if err != nil {
			return nil, err
		}
		ur.EndDate = &endDate
	}
	ur.Guests = dto.Guests

	return ur, nil
}

func ToReservationQueryFromDTO(dto *dto.ReservationQueryDTO) (*model.ReservationQuery, error) {
	rq := &model.ReservationQuery{
		Limit: dto.GetLimit(),
	}
	v := validator.New()

	if dto.StartDate != nil {
		startDate, err := v.ParseDate(*dto.StartDate)
		if err != nil {
			return nil, err
		}
		rq.StartDate = &startDate
	}
	if dto.EndDate != nil {
		endDate, err := v.ParseDate(*dto.EndDate)
		if err != nil {
			return nil, err
		}
		rq.EndDate = &endDate
	}

	if dto.ApartmentID != nil {
		apartmentID, err := uuid.Parse(*dto.ApartmentID)
		if err != nil {
			return nil, fmt.Errorf("invalid apartment id, must be UUID")
		}
		rq.ApartmentID = &apartmentID
	}

	return rq, nil
}
