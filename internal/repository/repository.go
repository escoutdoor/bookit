package repository

import (
	"context"
	"time"

	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
)

type ApartmentRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error)
	Create(ctx context.Context, in *model.CreateApartment, userID uuid.UUID) (*model.Apartment, error)
	Update(ctx context.Context, in *model.UpdateApartment, apartmentID uuid.UUID) (*model.Apartment, error)
	Delete(ctx context.Context, id uuid.UUID) error

	GetAll(ctx context.Context, in *model.ApartmentQuery) ([]*model.Apartment, error)
}

type ApartmentCache interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error)
	Set(ctx context.Context, in *model.Apartment) error
	Delete(ctx context.Context, id uuid.UUID) error
	Expire(ctx context.Context, id uuid.UUID) error
}

type CategoryRepository interface {
	Create(ctx context.Context, in *model.CreateCategory) (*model.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error)
	Update(ctx context.Context, in *model.UpdateCategory, categoryID uuid.UUID) (*model.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ReviewRepository interface {
	Create(ctx context.Context, in *model.CreateReview, renterID uuid.UUID) (*model.Review, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Review, error)

	GetAll(ctx context.Context, in *model.ReviewQuery) ([]*model.Review, error)
}

type UserRepository interface {
	Create(ctx context.Context, in *model.RegisterUser) (*model.User, error)
	Update(ctx context.Context, in *model.UpdateUser, userID uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}

type ReservationRepository interface {
	Create(ctx context.Context, in *model.CreateReservation, renterID uuid.UUID, total float64) (*model.Reservation, error)
	Update(ctx context.Context, in *model.UpdateReservation, rsrvID uuid.UUID) (*model.Reservation, error)
	Delete(ctx context.Context, reservationID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Reservation, error)

	GetAll(ctx context.Context, in *model.ReservationQuery) ([]*model.Reservation, error)
	GetBetweenDates(ctx context.Context, apartmentID uuid.UUID, startDate, endDate time.Time) ([]*model.Reservation, error)
	GetBetweenDatesExcludingID(ctx context.Context, rsrvID, apartmentID uuid.UUID, startDate, endDate time.Time) ([]*model.Reservation, error)
	IsRentedBefore(ctx context.Context, apatmentID, renterID uuid.UUID) (bool, error)
}

type AvatarRepository interface {
	Upload(ctx context.Context, in *model.UploadAvatar) (string, error)
}
