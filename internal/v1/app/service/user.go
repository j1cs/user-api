package service

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/domain"
)

type UserService struct {
	repository domain.UserRepository
	publisher  domain.UserPublisher
	logger     *zerolog.Logger
}

func NewUserService(repo domain.UserRepository, publisher domain.UserPublisher, logger *zerolog.Logger) *UserService {
	return &UserService{repo, publisher, logger}
}

func (u *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering get all service")
	return u.repository.FindAll(ctx)
}

func (u *UserService) GetOne(ctx context.Context, uuid uuid.UUID) (*domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering get one service")
	return u.repository.FindOne(ctx, uuid)
}

func (u *UserService) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering create service")
	saved, err := u.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	msgId, err := u.publisher.Publish(ctx, *saved, domain.NewHeader())
	u.logger.Info().
		Str("requestId", middleware.GetReqID(ctx)).
		Str("messageId", msgId).
		Msg("User published")
	if err != nil {
		return nil, err
	}
	return saved, nil
}

func (u *UserService) Update(ctx context.Context, uuid uuid.UUID, user domain.User) error {
	return u.repository.Update(ctx, uuid, user)
}

func (u *UserService) Delete(ctx context.Context, uuid uuid.UUID) error {
	return u.repository.Delete(ctx, uuid)
}
