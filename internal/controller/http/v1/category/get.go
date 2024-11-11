package category

import (
	"context"
	"errors"
	"net/http"

	httphelper "github.com/escoutdoor/bookit/internal/controller/http"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/model"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
)

func (c *controller) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	ctgr, err := c.uc.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, categoryuc.ErrNotFound) {
			resp.JSON(w, http.StatusNotFound, "category not found")
			return
		}

		c.log.Error("failed to get category by id", "error", err)
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
