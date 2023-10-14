package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindOne(ctx context.Context, uuid uuid.UUID) (*User, error)
	Save(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, uuid uuid.UUID, user User) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
