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
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
)

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	u, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	var dto dto.CreateApartmentDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToCreateApartmentFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	apart, err := c.uc.Create(ctx, in, u.ID)
	if err != nil {
		if errors.Is(err, categoryuc.ErrNotFound) {
			resp.Error(w, http.StatusNotFound, "category not found")
			return
		}

		c.log.Error("failed to create apartment", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Apartment *model.Apartment `json:"apartment"`
	}

	resp.JSON(w, http.StatusOK, response{
		Apartment: apart,
	})
}
