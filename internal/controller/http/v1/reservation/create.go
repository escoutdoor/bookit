package reservation

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httphelper "github.com/escoutdoor/bookit/internal/controller/http"
	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/converter"
	"github.com/escoutdoor/bookit/internal/model"
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	reservationuc "github.com/escoutdoor/bookit/internal/usecase/reservation"
)

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	uctx, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	var dto dto.CreateReservationDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToCreateReservationFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	rsrv, err := c.uc.Create(ctx, in, uctx.ID)
	if err != nil {
		switch {
		case errors.Is(err, apartmentuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "apartment not found")
			return
		case errors.Is(err, reservationuc.ErrAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "reservations for those days already exist")
			return
		default:
			c.log.Error("failed to create reservation", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	type response struct {
		Reservation *model.Reservation `json:"reservation"`
	}

	resp.JSON(w, http.StatusOK, response{
		Reservation: rsrv,
	})
}
