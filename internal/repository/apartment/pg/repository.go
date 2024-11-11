package pg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/model"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	"github.com/escoutdoor/bookit/internal/repository/apartment/pg/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/apartment/pg/model"
	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	"github.com/escoutdoor/bookit/internal/repository/code"
	userrepo "github.com/escoutdoor/bookit/internal/repository/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	db database.Client
}

func NewApartmentRepository(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentRepository.GetByID"
	query := "SELECT * FROM APARTMENTS WHERE ID = $1"

	var apartment repomodel.Apartment
	err := r.db.QueryRow(ctx, query, id).Scan(
		&apartment.ID,
		&apartment.Name,
		&apartment.Description,
		&apartment.Beds,
		&apartment.Bedrooms,
		&apartment.Bathrooms,
		&apartment.MaxGuests,
		&apartment.RentalPrice,
		&apartment.Latitude,
		&apartment.Longitude,
		&apartment.HostID,
		&apartment.CategoryID,
		&apartment.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, aparterr.ErrNotFound
		}
		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToApartmentFromRepository(&apartment), nil
}

func (r *repository) Create(ctx context.Context, in *model.CreateApartment, userID uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentRepository.Create"
	query := `
		INSERT INTO 
		APARTMENTS(
            NAME,
            DESCRIPTION,
            BEDS,
            BEDROOMS,
            BATHROOMS,
            MAX_GUESTS,
            RENTAL_PRICE,
            LATITUDE,
            LONGITUDE,
            HOST_ID,
            CATEGORY_ID
        )
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING *
	`
	args := []interface{}{
		in.Name,
		in.Description,
		in.Beds,
		in.Bedrooms,
		in.Bathrooms,
		in.MaxGuests,
		in.RentalPrice,
		in.Latitude,
		in.Longitude,
		userID,
		in.CategoryID,
	}

	var apartment repomodel.Apartment
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&apartment.ID,
		&apartment.Name,
		&apartment.Description,
		&apartment.Beds,
		&apartment.Bedrooms,
		&apartment.Bathrooms,
		&apartment.MaxGuests,
		&apartment.RentalPrice,
		&apartment.Latitude,
		&apartment.Longitude,
		&apartment.HostID,
		&apartment.CategoryID,
		&apartment.CreatedAt,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.ForeignKeyViolationCode {
			switch {
			case pgErr.ConstraintName == "apartments_category_id_fkey":
				return nil, categoryrepo.ErrNotFound
			case pgErr.ConstraintName == "apartments_user_id_fkey":
				return nil, userrepo.ErrNotFound
			}
		}
		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToApartmentFromRepository(&apartment), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdateApartment, apartmentID uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentRepository.Update"

	args := pgx.NamedArgs{}
	query := `
		UPDATE APARTMENTS
        SET 
	`

	var updates []string
	if in.Name != nil {
		updates = append(updates, "name=@name")
		args["name"] = in.Name
	}
	if in.Description != nil {
		updates = append(updates, "description=@description")
		args["description"] = in.Description
	}
	if in.Beds != nil {
		updates = append(updates, "beds=@beds")
		args["beds"] = in.Beds
	}
	if in.Bedrooms != nil {
		updates = append(updates, "bedrooms=@bedrooms")
		args["bedrooms"] = in.Bedrooms
	}
	if in.Bathrooms != nil {
		updates = append(updates, "bathrooms=@bathrooms")
		args["bathrooms"] = in.Bathrooms
	}
	if in.MaxGuests != nil {
		updates = append(updates, "max_guests=@max_guests")
		args["max_guests"] = in.MaxGuests
	}
	if in.RentalPrice != nil {
		updates = append(updates, "rental_price=@rental_price")
		args["rental_price"] = in.RentalPrice
	}
	if in.Latitude != nil {
		updates = append(updates, "latitude=@latitude")
		args["latitude"] = in.Latitude
	}
	if in.Longitude != nil {
		updates = append(updates, "longitude=@longitude")
		args["longitude"] = in.Longitude
	}
	if in.CategoryID != nil {
		updates = append(updates, "category_id=@category_id")
		args["category_id"] = in.CategoryID
	}

	if len(args) == 0 {
		return nil, aparterr.ErrNoFieldsToUpdate
	}

	query += fmt.Sprintf(" %s WHERE ID=@id RETURNING *", strings.Join(updates, ", "))
	args["id"] = apartmentID

	var apartment repomodel.Apartment
	err := r.db.QueryRow(ctx, query, args).Scan(
		&apartment.ID,
		&apartment.Name,
		&apartment.Description,
		&apartment.Beds,
		&apartment.Bedrooms,
		&apartment.Bathrooms,
		&apartment.MaxGuests,
		&apartment.RentalPrice,
		&apartment.Latitude,
		&apartment.Longitude,
		&apartment.HostID,
		&apartment.CategoryID,
		&apartment.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.ForeignKeyViolationCode {
			if pgErr.ConstraintName == "apartments_category_id_fkey" {
				return nil, aparterr.ErrNotFound
			}
		}

		return nil, fmt.Errorf("%s: query row: %s", op, err)
	}

	return converter.ToApartmentFromRepository(&apartment), nil
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "ApartmentRepository.Delete"
	query := `
        DELETE FROM APARTMENTS WHERE ID = $1
    `

	res, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: query row: %s", op, err)
	}

	if n := res.RowsAffected(); n == 0 {
		return fmt.Errorf("%s: couldn't delete apartment error: %s", op, err)
	}
	return nil
}

func (r *repository) GetAll(ctx context.Context, in *model.ApartmentQuery) ([]*model.Apartment, error) {
	const op = "ApartmentRepository.GetAll"

	args := pgx.NamedArgs{
		"limit": in.Limit,
	}
	query := "SELECT * FROM APARTMENTS WHERE 1=1"

	if in.Name != nil {
		query += " AND name=@name"
		args["name"] = in.Name
	}
	if in.Description != nil {
		query += " AND description=@description"
		args["description"] = in.Description
	}
	if in.Beds != nil {
		query += " AND beds=@beds"
		args["beds"] = in.Beds
	}
	if in.Bedrooms != nil {
		query += " AND bedrooms=@bedrooms"
		args["bedrooms"] = in.Bedrooms
	}
	if in.Bathrooms != nil {
		query += " AND bathrooms=@bathrooms"
		args["bathrooms"] = in.Bathrooms
	}
	if in.MaxGuests != nil {
		query += " AND max_guests=@max_guests"
		args["max_guests"] = in.MaxGuests
	}
	if in.CategoryID != nil {
		query += " AND category_id=@category_id"
		args["category_id"] = in.CategoryID
	}
	if in.MinRentalPrice != nil {
		query += " AND rental_price > @min_rental_price"
		args["min_rental_price"] = in.MinRentalPrice
	}
	if in.MaxRentalPrice != nil {
		query += " AND rental_price < @max_rental_price"
		args["max_rental_price"] = in.MaxRentalPrice
	}

	query += fmt.Sprintf(" ORDER BY %s LIMIT @limit", in.SortBy)

	rows, err := r.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %s", op, err)
	}
	defer rows.Close()

	aparts, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*repomodel.Apartment, error) {
		var apart repomodel.Apartment
		err := row.Scan(
			&apart.ID,
			&apart.Name,
			&apart.Description,
			&apart.Beds,
			&apart.Bedrooms,
			&apart.Bathrooms,
			&apart.MaxGuests,
			&apart.RentalPrice,
			&apart.Latitude,
			&apart.Longitude,
			&apart.HostID,
			&apart.CategoryID,
			&apart.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &apart, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: collect rows: %s", op, err)
	}

	return converter.ToApartmentsFromRepository(aparts), nil
}
