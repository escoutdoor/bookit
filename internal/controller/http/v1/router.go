package v1

import (
	"log/slog"
	"net/http"

	"github.com/escoutdoor/bookit/internal/config"
	"github.com/escoutdoor/bookit/internal/controller/http/middleware"
	"github.com/escoutdoor/bookit/internal/controller/http/resp"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/apartment"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/auth"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/avatar"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/category"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/reservation"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/review"
	"github.com/escoutdoor/bookit/internal/controller/http/v1/user"
	"github.com/escoutdoor/bookit/internal/usecase"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(
	h *chi.Mux,
	cfg *config.AppConfig,
	log *slog.Logger,
	authUC usecase.AuthUseCase,
	apartUC usecase.ApartmentUseCase,
	categoryUC usecase.CategoryUseCase,
	reviewUC usecase.ReviewUseCase,
	rsrvUC usecase.ReservationUseCase,
	userUC usecase.UserUseCase,
	avatarUC usecase.AvatarUseCase,
) {
	h.Use(chimiddle.StripSlashes)
	h.Use(chimiddle.Recoverer)
	h.Use(chimiddle.RequestID)

	authmiddle := middleware.NewAuthMiddleware(authUC, userUC, log)

	apartment := apartment.NewApartmentController(apartUC, log)
	rsrv := reservation.NewReservationController(rsrvUC, log)
	auth := auth.NewAuthController(authUC, log)
	category := category.NewCategoryController(categoryUC, log)
	review := review.NewReviewController(reviewUC, log)
	user := user.NewUserController(userUC, log)
	avatar := avatar.NewAvatarController(avatarUC, log)

	h.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		resp.JSON(w, http.StatusOK, map[string]interface{}{
			"env":     cfg.Env,
			"name":    cfg.Name,
			"version": cfg.Version,
		})
	})

	h.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", auth.Register)
			r.Post("/login", auth.Login)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", user.GetByID)

			r.With(authmiddle.Auth).Patch("/", user.Update)
		})

		r.Route("/avatars", func(r chi.Router) {
			r.Use(authmiddle.Auth)

			r.Post("/", avatar.Upload)
		})

		r.Route("/apartments", func(r chi.Router) {
			r.Get("/{id}", apartment.GetByID)
			r.Get("/", apartment.GetAll)

			r.Group(func(r chi.Router) {
				r.Use(authmiddle.Auth)

				r.Post("/", apartment.Create)
				r.Patch("/{id}", apartment.Update)
				r.Delete("/{id}", apartment.Delete)
			})
		})

		r.Route("/reservations", func(r chi.Router) {
			r.With(authmiddle.Auth).Post("/", rsrv.Create)
			r.With(authmiddle.Auth).Patch("/{id}", rsrv.Update)
			r.With(authmiddle.Auth).Get("/", rsrv.GetAll)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Get("/{id}", category.GetByID)

			// TODO: only for admin
			r.Group(func(r chi.Router) {
				r.Use(authmiddle.Auth)
				r.Use(authmiddle.IsAdmin)

				r.Post("/", category.Create)
				r.Patch("/{id}", category.Update)
				r.Delete("/{id}", category.Delete)
			})
		})

		r.Route("/reviews", func(r chi.Router) {
			r.With(authmiddle.Auth).Post("/", review.Create)

			r.Get("/", review.GetAll)
			r.Get("/{id}", review.GetByID)
		})
	})
}
