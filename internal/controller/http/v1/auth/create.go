package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/converter"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
)

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	var dto dto.RegisterUserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := dto.Validate(); len(err) > 0 {
		resp.JSON(w, http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	res, err := c.uc.Register(ctx, converter.ToRegisterUserFromDTO(&dto))
	if err != nil {
		switch {
		case errors.Is(err, authuc.ErrEmailAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "user with this email already exists")
			return
		case errors.Is(err, authuc.ErrPhoneNumberAlreadyExists):
			resp.Error(w, http.StatusBadRequest, "user with this phone number already exists")
			return
		default:
			c.log.Error("failed to register user", "error", err)
			resp.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	resp.JSON(w, http.StatusOK, res)
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		resp.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := dto.Validate(); len(err) > 0 {
		resp.JSON(w, http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	res, err := c.uc.Login(ctx, converter.ToLoginFromDTO(&dto))
	if err != nil {
		if errors.Is(err, authuc.ErrInvalidEmailOrPassword) {
			resp.Error(w, http.StatusBadRequest, "invalid email or password")
			return
		}

		c.log.Error("failed to login", "error", err)
		resp.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp.JSON(w, http.StatusOK, res)
}
