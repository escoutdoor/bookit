package reservation

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/converter"
	"github.com/escoutdoor/bookit/internal/model"
)

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	dto, err := c.getQueryParams(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToReservationQueryFromDTO(dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	rsrvs, err := c.uc.GetAll(ctx, in)
	if err != nil {
		c.log.Error("failed to get all reservation", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Reservations []*model.Reservation `json:"reservations"`
	}

	resp.JSON(w, http.StatusOK, response{
		Reservations: rsrvs,
	})
}

func (c *controller) getQueryParams(r *http.Request) (*dto.ReservationQueryDTO, error) {
	dto := &dto.ReservationQueryDTO{}

	apartmentID := r.URL.Query().Get("apartment_id")
	if len(apartmentID) > 0 {
		dto.ApartmentID = &apartmentID
	}
	startDate := r.URL.Query().Get("start_date")
	if len(startDate) > 0 {
		dto.StartDate = &startDate
	}
	endDate := r.URL.Query().Get("end_date")
	if len(endDate) > 0 {
		dto.EndDate = &endDate
	}

	limitStr := r.URL.Query().Get("limit")
	if len(limitStr) > 0 {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, fmt.Errorf("limit should be an integer")
		}
		dto.Limit = limit
	}

	return dto, nil
}
