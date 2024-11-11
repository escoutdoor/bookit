package middleware

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/internal/usecase"
	useruc "github.com/escoutdoor/bookit/internal/usecase/user"
)

const userCtxKey = "user"

type authMiddleware struct {
	log    *slog.Logger
	authUC usecase.AuthUseCase
	userUC usecase.UserUseCase
}

func NewAuthMiddleware(
	authUC usecase.AuthUseCase,
	userUC usecase.UserUseCase,
	log *slog.Logger,
) *authMiddleware {
	return &authMiddleware{
		log:    log,
		authUC: authUC,
		userUC: userUC,
	}
}

func (m *authMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "AuthMiddleware.Auth"
		jwtToken := r.Header.Get("Authorization")
		if len(jwtToken) == 0 {
			resp.Error(w, http.StatusUnauthorized, "invalid authorization token")
			return
		}
		jwtToken = jwtToken[len("Bearer "):]
		id, err := m.authUC.ParseToken(jwtToken)
		if err != nil {
			resp.Error(w, http.StatusUnauthorized, fmt.Sprintf("failed to parse token: %s", err))
			return
		}

		user, err := m.userUC.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, useruc.ErrNotFound) {
				resp.Error(w, http.StatusUnauthorized, "user not found")
				return
			}

			m.log.Error("failed to get user by id", "error", fmt.Sprintf("%s: %s", op, err))
			resp.Error(w, http.StatusInternalServerError, "failed to authorize")
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, user)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (m *authMiddleware) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "AuthMiddleware.IsAdmin"

		user, ok := r.Context().Value("user").(*model.User)
		if !ok {
			m.log.Error("failed to get user from ctx", "error", fmt.Sprintf("%s: %s", op, "failed to retrieve user from context"))
			resp.Error(w, http.StatusInternalServerError, "user not authorized")
			return
		}

		if user.Role != "admin" {
			resp.Error(w, http.StatusForbidden, "forbidden")
			return
		}
		next.ServeHTTP(w, r)
	})
}
