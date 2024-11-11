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
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	reservationuc "github.com/escoutdoor/bookit/internal/usecase/reservation"
)

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	uctx, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	rsrvID, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto dto.UpdateReservationDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToUpdateReservationFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	rsrv, err := c.uc.Update(ctx, in, rsrvID, uctx.ID)
	if err != nil {
		switch {
		case errors.Is(err, authuc.ErrForbidden):
			resp.Error(w, http.StatusForbidden, "forbidden")
			return
		case errors.Is(err, reservationuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "reservation not found")
			return
		case errors.Is(err, reservationuc.ErrAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "reservations for those days already exist")
			return
		case errors.Is(err, reservationuc.ErrNoFieldsToUpdate):
			resp.Error(w, http.StatusBadRequest, "no fields to update for reservation")
			return
		case errors.Is(err, reservationuc.ErrInvalidReservationRange):
			resp.Error(w, http.StatusBadRequest, "the end date must be after the start date")
			return
		default:
			c.log.Error("failed to update reservation", "error", err)
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
