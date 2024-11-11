package category

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/internal/repository/category/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/category/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	db database.Client
}

func NewCategoryRepository(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.CreateCategory) (*model.Category, error) {
	const op = "CategoryRepository.Create"
	query := `
		INSERT INTO CATEGORIES(NAME)
		VALUES($1)
		RETURNING *
	`

	args := []interface{}{
		in.Name,
	}

	var c repomodel.Category
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&c.ID,
		&c.Name,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToCategoryFromRepository(&c), nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	const op = "CategoryRepository.GetByID"
	query := `
		SELECT * FROM CATEGORIES WHERE ID = $1
	`

	var c repomodel.Category
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID,
		&c.Name,
		&c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToCategoryFromRepository(&c), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdateCategory, categoryID uuid.UUID) (*model.Category, error) {
	const op = "CategoryRepository.Update"

	args := pgx.NamedArgs{}
	query := `
		UPDATE CATEGORIES
		SET
	`

	var updates []string
	if in.Name != nil {
		updates = append(updates, "name=@name")
		args["name"] = in.Name
	}
	if len(args) == 0 {
		return nil, ErrNoFieldsToUpdate
	}

	query += fmt.Sprintf(" %s WHERE ID = @id RETURNING *", strings.Join(updates, ", "))
	args["id"] = categoryID

	var c repomodel.Category
	err := r.db.QueryRow(ctx, query, args).Scan(
		&c.ID,
		&c.Name,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToCategoryFromRepository(&c), nil
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "CategoryRepository.Delete"
	query := `
		DELETE FROM CATEGORIES WHERE ID = $1
	`

	res, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
