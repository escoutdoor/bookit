package review

import (
	"context"
	"errors"
	"fmt"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	"github.com/escoutdoor/bookit/internal/repository/code"
	"github.com/escoutdoor/bookit/internal/repository/review/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/review/model"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	db database.Client
}

func NewReviewRepository(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.CreateReview, renterID uuid.UUID) (*model.Review, error) {
	const op = "ReviewRepository.Create"

	query := `
        INSERT INTO REVIEWS(CONTENT, RATING, RENTER_ID, APARTMENT_ID)
        VALUES($1, $2, $3, $4)
        RETURNING *
    `
	args := []interface{}{
		in.Content,
		in.Rating,
		renterID,
		in.ApartmentID,
	}

	var rev repomodel.Review
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&rev.ID,
		&rev.Content,
		&rev.Rating,
		&rev.RenterID,
		&rev.ApartmentID,
		&rev.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.ForeignKeyViolationCode {
			switch {
			case pgErr.ConstraintName == "reviews_apartment_id_fkey":
				return nil, aparterr.ErrNotFound
			case pgErr.ConstraintName == "reviews_user_id_fkey":
				return nil, userrepo.ErrNotFound
			}
		}
		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToReviewFromRepository(&rev), nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Review, error) {
	const op = "ReviewRepository.GetByID"
	query := "SELECT * FROM REVIEWS WHERE ID = $1"

	var rev repomodel.Review
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rev.ID,
		&rev.Content,
		&rev.Rating,
		&rev.RenterID,
		&rev.ApartmentID,
		&rev.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToReviewFromRepository(&rev), nil
}

func (r *repository) GetAll(ctx context.Context, in *model.ReviewQuery) ([]*model.Review, error) {
	const op = "ReviewRepository.GetAll"

	args := pgx.NamedArgs{
		"limit": in.Limit,
	}
	query := "SELECT * FROM REVIEWS WHERE 1=1"

	if in.Rating != nil {
		query += " AND rating=@rating"
		args["rating"] = in.Rating
	}
	if in.RenterID != nil {
		query += " AND renter_id=@renter_id"
		args["renter_id"] = in.RenterID
	}
	if in.ApartmentID != nil {
		query += " AND apartment_id=@apartment_id"
		args["apartment_id"] = in.ApartmentID
	}

	query += " LIMIT @limit"

	rows, err := r.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %s", op, err)
	}
	defer rows.Close()

	revs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*repomodel.Review, error) {
		var rev repomodel.Review
		err := row.Scan(
			&rev.ID,
			&rev.Content,
			&rev.Rating,
			&rev.RenterID,
			&rev.ApartmentID,
			&rev.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &rev, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: collect rows: %s", op, err)
	}

	return converter.ToReviewsFromRepository(revs), nil
}
