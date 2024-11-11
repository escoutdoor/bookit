package pg

import (
	"fmt"

	"github.com/escoutdoor/bookit/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

func NewDB(ctx context.Context, cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dsn := generateConnStr(cfg)
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %s", err)
	}
	return db, nil
}

func generateConnStr(cfg *config.DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}
