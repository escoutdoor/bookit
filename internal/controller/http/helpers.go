package http

import (
	"fmt"
	"net/http"

	"github.com/escoutdoor/bookit/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetIDParam(r *http.Request) (uuid.UUID, error) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		return uuid.Nil, fmt.Errorf("id parameter is required")
	}

	pid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid id parameter, should be uuid")
	}
	return pid, nil
}

func GetUserFromCtx(r *http.Request) (*model.User, error) {
	u, ok := r.Context().Value("user").(*model.User)
	if !ok {
		return nil, fmt.Errorf("couldn't get user from context")
	}
	return u, nil
}
