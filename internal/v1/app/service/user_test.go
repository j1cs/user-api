//go:build service

package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/j1cs/api-user/internal/v1/domain"
	"github.com/j1cs/api-user/internal/v1/domain/mocks"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var (
	testLogger *zerolog.Logger
)

func TestNewUserService(t *testing.T) {
	userRepo := new(mocks.UserRepository)
	userPublisher := new(mocks.UserPublisher)
	testLogger = log.InitializeLogger(zerolog.InfoLevel)
	userService := NewUserService(userRepo, userPublisher, testLogger)
	assert.NotNil(t, userService)

	ctx := context.Background()
	sampleUsers := []domain.User{
		{Email: "test1@example.com", Name: "Test User 1"},
		{Email: "test2@example.com", Name: "Test User 2"},
	}
	sampleUUID := uuid.New()

	t.Run("GetAll", func(t *testing.T) {
		userRepo.On("FindAll", ctx).Return(sampleUsers, nil).Once()
		users, err := userService.GetAll(ctx)
		assert.NoError(t, err)
		assert.Equal(t, sampleUsers, users)
		userRepo.AssertExpectations(t)
	})

	t.Run("GetOne", func(t *testing.T) {
		userRepo.On("FindOne", ctx, mock.AnythingOfType("uuid.UUID")).Return(&sampleUsers[0], nil).Once()
		user, err := userService.GetOne(ctx, sampleUUID)
		assert.NoError(t, err)
		assert.Equal(t, &sampleUsers[0], user)
		userRepo.AssertExpectations(t)
	})

	t.Run("Create", func(t *testing.T) {
		userRepo.On("Save", ctx, mock.AnythingOfType("domain.User")).Return(&sampleUsers[0], nil).Once()
		userPublisher.On("Publish", ctx, mock.AnythingOfType("domain.User"), mock.AnythingOfType("domain.Header")).Return("msgID", nil).Once()

		user, err := userService.Create(ctx, sampleUsers[0])
		assert.NoError(t, err)
		assert.Equal(t, &sampleUsers[0], user)
		userRepo.AssertExpectations(t)
		userPublisher.AssertExpectations(t)
	})

	t.Run("Update", func(t *testing.T) {
		userRepo.On("Update", ctx, sampleUUID, sampleUsers[0]).Return(nil).Once()
		err := userService.Update(ctx, sampleUUID, sampleUsers[0])
		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		userRepo.On("Delete", ctx, sampleUUID).Return(nil).Once()
		err := userService.Delete(ctx, sampleUUID)
		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("CreateFail", func(t *testing.T) {
		userRepo.On("Save", ctx, mock.AnythingOfType("domain.User")).Return(nil, errors.New("some error")).Once()
		_, err := userService.Create(ctx, sampleUsers[0])
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
	})
}
