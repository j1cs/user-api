package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, uuid uuid.UUID) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, uuid uuid.UUID, user User) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
