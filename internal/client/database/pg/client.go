package pg

import (
	"context"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.DBConfig) (database.Client, error) {
	db, err := NewDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &pgClient{
		db: db,
	}, nil
}

func (cl *pgClient) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return cl.db.QueryRow(ctx, sql, args...)
}

func (cl *pgClient) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return cl.db.Query(ctx, sql, args...)
}

func (cl *pgClient) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return cl.Exec(ctx, sql, arguments...)
}
