package user

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

	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	useruc "github.com/escoutdoor/bookit/internal/usecase/user"
)

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	uctx, err := httphelper.GetUserFromCtx(r)
	if err != nil {
		resp.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	var dto dto.UpdateUserDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if errs := dto.Validate(); len(errs) > 0 {
		resp.JSON(w, http.StatusBadRequest, errs)
		return
	}

	in, err := converter.ToUpdateUserFromDTO(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	uu, err := c.uc.Update(ctx, in, uctx.ID)
	if err != nil {
		switch {
		case errors.Is(err, authuc.ErrEmailAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "user with this email address already exists")
			return
		case errors.Is(err, authuc.ErrPhoneNumberAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "user with this phone number already exists")
			return
		case errors.Is(err, useruc.ErrNoFieldsToUpdate):
			resp.Error(w, http.StatusBadRequest, "no fields to update for user")
			return
		default:
			c.log.Error("failed to update user", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	type response struct {
		User *model.User `json:"user"`
	}

	resp.JSON(w, http.StatusOK, response{
		User: uu,
	})
}
