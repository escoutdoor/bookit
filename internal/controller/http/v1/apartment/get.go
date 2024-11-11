package apartment

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
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
)

func (c *controller) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	apart, err := c.uc.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, apartmentuc.ErrNotFound) {
			resp.Error(w, http.StatusNotFound, "apartment not found")
			return
		}

		c.log.Error("failed to get apartment by id", "error", err)
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

	in, err := converter.ToApartmentQueryFromDTO(dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	aparts, err := c.uc.GetAll(ctx, in)
	if err != nil {
		c.log.Error("failed to get all apartments", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Apartments []*model.Apartment `json:"apartments"`
	}

	resp.JSON(w, http.StatusOK, response{
		Apartments: aparts,
	})
}

func (c *controller) getQueryParams(r *http.Request) (*dto.ApartmentQueryDTO, error) {
	dto := &dto.ApartmentQueryDTO{}

	name := r.URL.Query().Get("name")
	if len(name) > 0 {
		dto.Name = &name
	}

	desc := r.URL.Query().Get("description")
	if len(desc) > 0 {
		dto.Description = &desc
	}

	bedsStr := r.URL.Query().Get("beds")
	if len(bedsStr) > 0 {
		beds, err := strconv.Atoi(bedsStr)
		if err != nil {
			return nil, fmt.Errorf("beds value should be an integer")
		}
		dto.Beds = &beds
	}

	bedroomsStr := r.URL.Query().Get("bedroms")
	if len(bedroomsStr) > 0 {
		bedrooms, err := strconv.Atoi(bedroomsStr)
		if err != nil {
			return nil, fmt.Errorf("bedrooms value should be an integer")
		}
		dto.Bedrooms = &bedrooms
	}

	bathroomsStr := r.URL.Query().Get("bathrooms")
	if len(bathroomsStr) > 0 {
		bathrooms, err := strconv.Atoi(bathroomsStr)
		if err != nil {
			return nil, fmt.Errorf("bathrooms value should be an integer")
		}
		dto.Bathrooms = &bathrooms
	}

	maxGuestsStr := r.URL.Query().Get("max_guests")
	if len(maxGuestsStr) > 0 {
		maxGuests, err := strconv.Atoi(maxGuestsStr)
		if err != nil {
			return nil, fmt.Errorf("max_guests value should be an integer")
		}
		dto.MaxGuests = &maxGuests
	}

	categoryID := r.URL.Query().Get("category_id")
	if len(categoryID) > 0 {
		dto.CategoryID = &categoryID
	}

	limitStr := r.URL.Query().Get("limit")
	if len(limitStr) > 0 {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, fmt.Errorf("limit should be an integer")
		}
		dto.Limit = limit
	}

	sortBy := r.URL.Query().Get("sort_by")
	if len(sortBy) > 0 {
		dto.SortBy = sortBy
	}

	return dto, nil
}
