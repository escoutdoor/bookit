package usecase

import (
	"context"
	"io"

	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
)

type ApartmentUseCase interface {
	Create(ctx context.Context, in *model.CreateApartment, hostID uuid.UUID) (*model.Apartment, error)
	Update(ctx context.Context, in *model.UpdateApartment, apartmentID, hostID uuid.UUID) (*model.Apartment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error)
	Delete(ctx context.Context, id, hostID uuid.UUID) error

	GetAll(ctx context.Context, in *model.ApartmentQuery) ([]*model.Apartment, error)
}

type ReservationUseCase interface {
	Create(ctx context.Context, in *model.CreateReservation, renterID uuid.UUID) (*model.Reservation, error)
	Update(ctx context.Context, in *model.UpdateReservation, rsrvID, renterID uuid.UUID) (*model.Reservation, error)
	GetAll(ctx context.Context, in *model.ReservationQuery) ([]*model.Reservation, error)
}

type CategoryUseCase interface {
	Create(ctx context.Context, in *model.CreateCategory) (*model.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error)
	Update(ctx context.Context, in *model.UpdateCategory, categoryID uuid.UUID) (*model.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ReviewUseCase interface {
	Create(ctx context.Context, in *model.CreateReview, renterID uuid.UUID) (*model.Review, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Review, error)
	GetAll(ctx context.Context, in *model.ReviewQuery) ([]*model.Review, error)
}

type AuthUseCase interface {
	Register(ctx context.Context, in *model.RegisterUser) (*model.AuthResponse, error)
	Login(ctx context.Context, in *model.Login) (*model.AuthResponse, error)
	ParseToken(jwtToken string) (uuid.UUID, error)
}

type UserUseCase interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Update(ctx context.Context, in *model.UpdateUser, userID uuid.UUID) (*model.User, error)
}

type AvatarUseCase interface {
	Upload(ctx context.Context, payload io.Reader, size int64) (string, error)
}
