package reservation

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/internal/repository/code"
	"github.com/escoutdoor/bookit/internal/repository/reservation/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/reservation/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	db database.Client
}

func NewReservationRepository(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Reservation, error) {
	const op = "ReservationRepository.GetByID"
	query := "SELECT * FROM RESERVATIONS WHERE ID = $1"

	var rsrv repomodel.Reservation
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rsrv.ID,
		&rsrv.RenterID,
		&rsrv.ApartmentID,
		&rsrv.StartDate,
		&rsrv.EndDate,
		&rsrv.Guests,
		&rsrv.Total,
		&rsrv.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToReservationFromRepository(&rsrv), nil
}

func (r *repository) Create(
	ctx context.Context,
	in *model.CreateReservation,
	renterID uuid.UUID,
	total float64,
) (*model.Reservation, error) {
	const op = "ReservationRepository.Create"
	query := `
        INSERT INTO RESERVATIONS(RENTER_ID, APARTMENT_ID, START_DATE, END_DATE, GUESTS, TOTAL)
        VALUES($1, $2, $3, $4, $5, $6)
        RETURNING *
    `

	args := []interface{}{
		renterID,
		in.ApartmentID,
		in.StartDate,
		in.EndDate,
		in.Guests,
		total,
	}
	var rsrv repomodel.Reservation
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&rsrv.ID,
		&rsrv.RenterID,
		&rsrv.ApartmentID,
		&rsrv.StartDate,
		&rsrv.EndDate,
		&rsrv.Guests,
		&rsrv.Total,
		&rsrv.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToReservationFromRepository(&rsrv), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdateReservation, rsrvID uuid.UUID) (*model.Reservation, error) {
	const op = "ReservationRepository.Update"

	args := pgx.NamedArgs{}
	query := `
        UPDATE RESERVATIONS
        SET 
    `

	var updates []string
	if in.StartDate != nil {
		updates = append(updates, "start_date=@start_date")
		args["start_date"] = in.StartDate
	}
	if in.EndDate != nil {
		updates = append(updates, "end_date=@end_date")
		args["end_date"] = in.EndDate
	}
	if in.Guests != nil {
		updates = append(updates, "guests=@guests")
		args["guests"] = in.Guests
	}
	if in.Total != nil {
		updates = append(updates, "total=@total")
		args["total"] = in.Total
	}

	if len(updates) == 0 {
		return nil, ErrNoFieldsToUpdate
	}

	query += fmt.Sprintf(" %s WHERE ID=@id RETURNING *", strings.Join(updates, ", "))
	args["id"] = rsrvID

	var rsrv repomodel.Reservation
	err := r.db.QueryRow(ctx, query, args).Scan(
		&rsrv.ID,
		&rsrv.RenterID,
		&rsrv.ApartmentID,
		&rsrv.StartDate,
		&rsrv.EndDate,
		&rsrv.Guests,
		&rsrv.Total,
		&rsrv.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToReservationFromRepository(&rsrv), nil
}

func (r *repository) GetAll(ctx context.Context, in *model.ReservationQuery) ([]*model.Reservation, error) {
	const op = "ReservationRepository.GetAll"

	args := pgx.NamedArgs{
		"limit": in.Limit,
	}
	query := "SELECT * FROM RESERVATIONS WHERE 1=1"

	if in.StartDate != nil {
		query += " AND start_date>=@start_date"
		args["start_date"] = in.StartDate
	}
	if in.EndDate != nil {
		query += " AND end_date<=@end_date"
		args["end_date"] = in.EndDate
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

	rsrvs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*repomodel.Reservation, error) {
		var rsrv repomodel.Reservation
		err := row.Scan(
			&rsrv.ID,
			&rsrv.RenterID,
			&rsrv.ApartmentID,
			&rsrv.StartDate,
			&rsrv.EndDate,
			&rsrv.Guests,
			&rsrv.Total,
			&rsrv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &rsrv, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: collect rows: %s", op, err)
	}

	return converter.ToReservationsFromRepository(rsrvs), nil
}

func (r *repository) GetBetweenDates(ctx context.Context, apartmentID uuid.UUID, startDate, endDate time.Time) ([]*model.Reservation, error) {
	const op = "ReservationRepository.GetBetweenDates"
	query := `
        SELECT * FROM RESERVATIONS WHERE APARTMENT_ID = $1 
        AND (START_DATE < $2 AND END_DATE > $3)
    `

	args := []interface{}{
		apartmentID,
		endDate,
		startDate,
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %s", op, err)
	}
	defer rows.Close()

	rsrvs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*repomodel.Reservation, error) {
		var rsrv repomodel.Reservation
		err := row.Scan(
			&rsrv.ID,
			&rsrv.RenterID,
			&rsrv.ApartmentID,
			&rsrv.StartDate,
			&rsrv.EndDate,
			&rsrv.Guests,
			&rsrv.Total,
			&rsrv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &rsrv, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: collect rows: %s", op, err)
	}

	return converter.ToReservationsFromRepository(rsrvs), nil
}

func (r *repository) GetBetweenDatesExcludingID(ctx context.Context, apartmentID, rsrvID uuid.UUID, startDate, endDate time.Time) ([]*model.Reservation, error) {
	const op = "ReservationRepository.GetBetweenDates"
	query := `
        SELECT * FROM RESERVATIONS WHERE APARTMENT_ID = $1 
        AND (START_DATE < $2 AND END_DATE > $3) AND ID != $4
    `

	args := []interface{}{
		apartmentID,
		endDate,
		startDate,
		rsrvID,
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %s", op, err)
	}
	defer rows.Close()

	rsrvs, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*repomodel.Reservation, error) {
		var rsrv repomodel.Reservation
		err := row.Scan(
			&rsrv.ID,
			&rsrv.RenterID,
			&rsrv.ApartmentID,
			&rsrv.StartDate,
			&rsrv.EndDate,
			&rsrv.Guests,
			&rsrv.Total,
			&rsrv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &rsrv, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: collect rows: %s", op, err)
	}

	return converter.ToReservationsFromRepository(rsrvs), nil
}

func (r *repository) Delete(ctx context.Context, reservationID uuid.UUID) error {
	const op = "ReservationRepository.Delete"
	query := "DELETE FROM RESERVATIONS WHERE ID = $1"

	res, err := r.db.Exec(ctx, query, reservationID)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%s: couldn't delete reservation error: %s", op, err)
	}
	return nil
}

func (r *repository) IsRentedBefore(ctx context.Context, apartmentID, renterID uuid.UUID) (bool, error) {
	const op = "ReservationRepository.IsRentedBefore"
	query := `
        SELECT id FROM RESERVATIONS WHERE APARTMENT_ID = $1 AND RENTER_ID = $2 
        LIMIT 1
    `

	var id uuid.NullUUID
	err := r.db.QueryRow(ctx, query, apartmentID, renterID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.ForeignKeyViolationCode {
			switch {
			case pgErr.ConstraintName == "reservations_apartments_id_fkey":
				return false, ErrNotFound
			}
		}
		return false, fmt.Errorf("%s: query row: %s", op, err)
	}

	return id.Valid, nil
}
