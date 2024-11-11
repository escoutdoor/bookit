package review

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	httphelper "github.com/escoutdoor/bookit/internal/controller/http"
	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/converter"
	"github.com/escoutdoor/bookit/internal/model"
	reviewuc "github.com/escoutdoor/bookit/internal/usecase/review"
)

func (c *controller) GetByID(w http.ResponseWriter, r *http.Request) {
	revID, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.Background()
	rev, err := c.uc.GetByID(ctx, revID)
	if err != nil {
		if errors.Is(err, reviewuc.ErrNotFound) {
			resp.Error(w, http.StatusNotFound, "review not found")
			return
		}

		c.log.Error("failed to get review by id", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Review *model.Review `json:"review"`
	}

	resp.JSON(w, http.StatusOK, response{
		Review: rev,
	})
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	dto, err := c.getQueryParams(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	in, err := converter.ToReviewQueryFromDTO(dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	revs, err := c.uc.GetAll(ctx, in)
	if err != nil {
		c.log.Error("failed to get reviews", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Reviews []*model.Review `json:"reviews"`
	}

	resp.JSON(w, http.StatusOK, response{
		Reviews: revs,
	})
}

func (c *controller) getQueryParams(r *http.Request) (*dto.ReviewQueryDTO, error) {
	dto := &dto.ReviewQueryDTO{}

	ratingStr := r.URL.Query().Get("rating")
	if len(ratingStr) > 0 {
		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			return nil, fmt.Errorf("rating should be an integer")
		}
		dto.Rating = &rating
	}

	renterID := r.URL.Query().Get("renter_id")
	if len(renterID) > 0 {
		dto.RenterID = &renterID
	}

	apartmentID := r.URL.Query().Get("apartment_id")
	if len(apartmentID) > 0 {
		dto.ApartmentID = &apartmentID
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
