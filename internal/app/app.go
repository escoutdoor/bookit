package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	rediscl "github.com/escoutdoor/bookit/internal/client/cache/redis"
	"github.com/escoutdoor/bookit/internal/client/database/pg"
	miniocl "github.com/escoutdoor/bookit/internal/client/s3/minio"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/escoutdoor/bookit/internal/controller/http/v1"
	apartpgrepo "github.com/escoutdoor/bookit/internal/repository/apartment/pg"
	apartredis "github.com/escoutdoor/bookit/internal/repository/apartment/redis"
	avatarrepo "github.com/escoutdoor/bookit/internal/repository/avatar"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	reservationrepo "github.com/escoutdoor/bookit/internal/repository/reservation"
	reviewrepo "github.com/escoutdoor/bookit/internal/repository/review"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"

	apartmentuc "github.com/escoutdoor/bookit/internal/usecase/apartment"
	authuc "github.com/escoutdoor/bookit/internal/usecase/auth"
	avataruc "github.com/escoutdoor/bookit/internal/usecase/avatar"
	categoryuc "github.com/escoutdoor/bookit/internal/usecase/category"
	reservationuc "github.com/escoutdoor/bookit/internal/usecase/reservation"
	reviewuc "github.com/escoutdoor/bookit/internal/usecase/review"
	useruc "github.com/escoutdoor/bookit/internal/usecase/user"
	"github.com/escoutdoor/bookit/pkg/httpserver"

	"github.com/go-chi/chi/v5"
)

// for testing
// type keyGen struct{}
//
// func (kg *keyGen) Generate(id uuid.UUID) string {
// 	return fmt.Sprintf("alomalo:%s", id)
// }

func Run(ctx context.Context, cfg *config.Config, log *slog.Logger) error {
	db, err := pg.New(ctx, &cfg.Database)
	if err != nil {
		return err
	}

	cacheClient, err := rediscl.NewClient(ctx, &cfg.Redis)
	if err != nil {
		return err
	}

	s3Client, err := miniocl.NewClient(ctx, &cfg.S3)
	if err != nil {
		return err
	}

	apartCache := apartredis.NewApartmentCache(
		cacheClient,
		&apartredis.DefaultKeyGenerator{},
		cfg.Redis.ApartmentTTL,
	)

	apartRepo := apartpgrepo.NewApartmentRepository(db)
	categoryRepo := categoryrepo.NewCategoryRepository(db)
	reviewRepo := reviewrepo.NewReviewRepository(db)
	rsrvRepo := reservationrepo.NewReservationRepository(db)
	userRepo := userrepo.NewUserRepository(db)
	avatarRepo := avatarrepo.NewAvatarRepository(s3Client)

	authUC := authuc.NewAuthUseCase(userRepo, cfg.JWT.SignKey, cfg.JWT.TokenTTL)
	apartUC := apartmentuc.NewApartmentUseCase(apartRepo, apartCache)
	categoryUC := categoryuc.NewCategoryUseCase(categoryRepo)
	reviewUC := reviewuc.NewReviewUseCase(reviewRepo, rsrvRepo)
	rsrvUC := reservationuc.NewReservationUseCase(rsrvRepo, apartRepo)
	userUC := useruc.NewUserUseCase(userRepo)
	avatarUC := avataruc.NewAvatarUseCase(avatarRepo)

	handler := chi.NewMux()
	v1.NewRouter(
		handler,
		&cfg.App,
		log,
		authUC,
		apartUC,
		categoryUC,
		reviewUC,
		rsrvUC,
		userUC,
		avatarUC,
	)

	httpSrv := httpserver.New(handler, &cfg.HTTPSrv)

	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		log.Info("server is running", "port", cfg.HTTPSrv.Port)
		err := httpSrv.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("failed to run server", "error", err)
		}
		close(quitch)
	}()

	<-quitch

	log.Info("server is shutting down..")
	httpSrv.GracefulStop()

	log.Info("httpserver successfully shut down")
	return nil
}
