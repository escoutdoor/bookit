package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/escoutdoor/bookit/internal/client/database"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/escoutdoor/bookit/internal/repository/code"
	"github.com/escoutdoor/bookit/internal/repository/user/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/user/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	db database.Client
}

func NewUserRepository(db database.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, in *model.RegisterUser) (*model.User, error) {
	const op = "UserRepository.Create"
	query := `
		INSERT INTO USERS(EMAIL, PASSWORD, FIRST_NAME, LAST_NAME, DATE_OF_BIRTH)
		VALUES($1, $2, $3, $4, $5)
		RETURNING *
	`

	args := []interface{}{
		in.Email,
		in.Password,
		in.FirstName,
		in.LastName,
		in.DOB,
	}

	var u repomodel.User
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.FirstName,
		&u.LastName,
		&u.DOB,
		&u.AvatarURL,
		&u.PhoneNumber,
		&u.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.UniqueViolationCode {
			switch {
			case pgErr.ConstraintName == "users_email_key":
				return nil, ErrEmailAlreadyExists
			case pgErr.ConstraintName == "users_phone_number_key":
				return nil, ErrPhoneNumberAlreadyExists
			}
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToUserFromRepository(&u), nil
}

func (r *repository) Update(ctx context.Context, in *model.UpdateUser, userID uuid.UUID) (*model.User, error) {
	const op = "UserRepository.Update"

	args := pgx.NamedArgs{}
	query := "UPDATE USERS SET"

	var updates []string
	if in.Email != nil {
		updates = append(updates, "email=@email")
		args["email"] = in.Email
	}
	if in.Password != nil {
		updates = append(updates, "password=@password")
		args["password"] = in.Password
	}
	if in.FirstName != nil {
		updates = append(updates, "first_name=@first_name")
		args["first_name"] = in.FirstName
	}
	if in.LastName != nil {
		updates = append(updates, "last_name=@last_name")
		args["last_name"] = in.LastName
	}
	if in.DOB != nil {
		updates = append(updates, "date_of_birth=@date_of_birth")
		args["date_of_birth"] = in.DOB
	}
	if in.AvatarURL != nil {
		updates = append(updates, "avatar_url=@avatar_url")
		args["avatar_url"] = in.AvatarURL
	}
	if in.PhoneNumber != nil {
		updates = append(updates, "phone_number=@phone_number")
		args["phone_number"] = in.PhoneNumber
	}
	if len(args) == 0 {
		return nil, ErrNoFieldsToUpdate
	}

	query += fmt.Sprintf(" %s WHERE ID=@id RETURNING *", strings.Join(updates, ", "))
	args["id"] = userID

	var u repomodel.User
	err := r.db.QueryRow(ctx, query, args).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.FirstName,
		&u.LastName,
		&u.DOB,
		&u.AvatarURL,
		&u.PhoneNumber,
		&u.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == code.UniqueViolationCode {
			switch {
			case pgErr.ConstraintName == "users_email_key":
				return nil, ErrEmailAlreadyExists
			case pgErr.ConstraintName == "users_phone_number_key":
				return nil, ErrPhoneNumberAlreadyExists
			}
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToUserFromRepository(&u), nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	const op = "UserRepository.GetByID"
	query := "SELECT * FROM USERS WHERE ID = $1"

	var u repomodel.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.FirstName,
		&u.LastName,
		&u.DOB,
		&u.AvatarURL,
		&u.PhoneNumber,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToUserFromRepository(&u), nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	const op = "UserRepository.GetByEmail"
	query := "SELECT * FROM USERS WHERE EMAIL = $1"

	var u repomodel.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.FirstName,
		&u.LastName,
		&u.DOB,
		&u.AvatarURL,
		&u.PhoneNumber,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return converter.ToUserFromRepository(&u), nil
}
