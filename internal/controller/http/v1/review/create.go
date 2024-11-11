package review

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
	reviewuc "github.com/escoutdoor/bookit/internal/usecase/review"
)

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	usr, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	var dto dto.CreateReviewDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errm := dto.Validate(); len(errm) > 0 {
		resp.JSON(w, http.StatusBadRequest, errm)
		return
	}

	in, err := converter.ToCreateReviewFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	rev, err := c.uc.Create(ctx, in, usr.ID)
	if err != nil {
		switch {
		case errors.Is(err, apartmentuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "apartment not found")
			return
		case errors.Is(err, reviewuc.ErrNotAllowedToCreateReview):
			resp.Error(w, http.StatusBadRequest, "you are not allowed to create reviews for this apartment")
			return
		default:
			c.log.Error("failed to create review", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	type response struct {
		Review *model.Review `json:"review"`
	}

	resp.JSON(w, http.StatusOK, response{
		Review: rev,
	})
}
