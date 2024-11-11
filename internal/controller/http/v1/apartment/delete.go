package apartment

import (
	"context"
	"errors"
	"net/http"

	httphelper "github.com/escoutdoor/bookit/internal/controller/http"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
)

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()
	err = c.uc.Delete(ctx, apartmentID, u.ID)
	if err != nil {
		switch {
		case errors.Is(err, apartmentuc.ErrNotFound):
			resp.Error(w, http.StatusNotFound, "apartment not found")
			return
		case errors.Is(err, authuc.ErrForbidden):
			resp.Error(w, http.StatusForbidden, "forbidden")
			return
		default:
			c.log.Error("failed to delete apartment", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	resp.JSON(w, http.StatusOK, map[string]string{
		"message": "apartment deleted",
	})
}
