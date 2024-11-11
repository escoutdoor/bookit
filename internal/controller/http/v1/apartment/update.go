package apartment

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
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
)

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	apartmentID, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	var dto dto.UpdateApartmentDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToUpdateApartmentFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	apart, err := c.uc.Update(ctx, in, apartmentID, u.ID)
	if err != nil {
		switch {
		case errors.Is(err, apartmentuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "apartment not found")
			return
		case errors.Is(err, categoryuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "category not found")
			return
		case errors.Is(err, apartmentuc.ErrNoFieldsToUpdate):
			resp.Error(w, http.StatusBadRequest, "no fields to update for apartment")
			return
		case errors.Is(err, authuc.ErrForbidden):
			resp.Error(w, http.StatusForbidden, "forbidden")
			return
		default:
			c.log.Error("failed to update apartment", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	type response struct {
		Apartment *model.Apartment `json:"apartment"`
	}

	resp.JSON(w, http.StatusOK, response{
		Apartment: apart,
	})
}
