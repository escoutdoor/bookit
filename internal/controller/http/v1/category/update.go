package category

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

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	categoryID, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto dto.UpdateCategoryDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errs := dto.Validate(); len(errs) > 0 {
		resp.JSON(w, http.StatusBadRequest, errs)
		return
	}

	ctx := context.Background()
	ctgr, err := c.uc.Update(ctx, converter.ToUpdateCategoryFromDTO(&dto), categoryID)
	if err != nil {
		switch {
		case errors.Is(err, categoryuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "category not found")
			return
		case errors.Is(err, categoryuc.ErrNoFieldsToUpdate):
			resp.Error(w, http.StatusBadRequest, "no fields to update for category")
			return
		default:
			c.log.Error("failed to update category", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	type response struct {
		Category *model.Category `json:"category"`
	}

	resp.JSON(w, http.StatusOK, response{
		Category: ctgr,
	})
}
