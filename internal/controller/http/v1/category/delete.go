package category

import (
	"context"
	"errors"
	"net/http"

	httphelper "github.com/escoutdoor/bookit/internal/controller/http"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
)

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := httphelper.GetIDParam(r)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	err = c.uc.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, categoryuc.ErrNotFound) {
			resp.Error(w, http.StatusBadRequest, "category not found")
			return
		}

		c.log.Error("failed to delete category", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp.JSON(w, http.StatusOK, map[string]string{
		"message": "category deleted",
	})
}
