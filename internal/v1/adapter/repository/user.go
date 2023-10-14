package repository

import (
	"context"
	"errors"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"github.com/j1cs/api-user/internal/v1/adapter/entity/db"
	"github.com/j1cs/api-user/internal/v1/domain"
)

type UserRepository struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewUserRepository(db *gorm.DB, logger *zerolog.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (u *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering find all repository")
	var entities []db.User

	result := u.db.WithContext(ctx).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	users, err := UserEntitiesToDomains(entities)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) FindOne(ctx context.Context, uuid uuid.UUID) (*domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering find one repository")
	var userEntity db.User
	if err := u.db.WithContext(ctx).Where("uuid = ?", uuid).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil if not found
		}
		return nil, err
	}
	user, err := userEntity.ToDomain()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Save(ctx context.Context, domainUser domain.User) (*domain.User, error) {
	u.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering create repository")
	var user db.User
	user.FromDomain(&domainUser)
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	domainUser.UUID = user.UUID

	return &domainUser, nil
}

func (u *UserRepository) Update(ctx context.Context, uuid uuid.UUID, user domain.User) error {
	return u.db.WithContext(ctx).Model(&db.User{}).Where("uuid = ?", uuid).Updates(user).Error
}

func (u *UserRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	return u.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&db.User{}).Error
}
