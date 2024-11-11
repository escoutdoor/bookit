package category

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/converter"
	"github.com/escoutdoor/bookit/internal/model"
)

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateCategoryDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errs := dto.Validate(); len(errs) > 0 {
		resp.JSON(w, http.StatusBadRequest, errs)
		return
	}

	ctx := context.Background()
	ctgr, err := c.uc.Create(ctx, converter.ToCreateCategoryFromDTO(&dto))
	if err != nil {
		c.log.Error("failed to create category", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Category *model.Category `json:"category"`
	}

	resp.JSON(w, http.StatusOK, response{
		Category: ctgr,
	})
}
