//go:build port

package port

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/j1cs/api-user/internal/v1/app"
	"github.com/j1cs/api-user/internal/v1/domain"
	"github.com/j1cs/api-user/internal/v1/domain/mocks"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var (
	server      *Server
	userService *mocks.UserService
	testLogger  *zerolog.Logger
	ctx         context.Context
	sampleUsers []domain.User
)

func TestMain(m *testing.M) {
	userService = new(mocks.UserService)
	testLogger = log.InitializeLogger(zerolog.InfoLevel)
	services := &app.Application{
		Service: &app.Service{User: userService},
	}
	server = NewServer(services, testLogger)
	ctx = context.Background()
	sampleUsers = []domain.User{
		{Email: "test1@example.com", Name: "Test User 1"},
		{Email: "test2@example.com", Name: "Test User 2"},
	}
	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	req := GetUsersRequestObject{}

	t.Run("testGetUsersSuccess", func(t *testing.T) {
		userService.On("GetAll", ctx).Return(sampleUsers, nil).Once()
		resp, _ := server.GetUsers(ctx, req)

		expectedResp := GetUsers200JSONResponse(UserDomainsToUsers(sampleUsers))
		assert.Equal(t, expectedResp, resp)

		userService.AssertExpectations(t)
	})
}

func TestPostUsers(t *testing.T) {
	req := PostUsersRequestObject{
		Body: &UserInput{
			Email: "test3@example.com",
			Name:  "Test User 3",
		},
	}

	t.Run("testPostUsersSuccess", func(t *testing.T) {
		domainUser := domain.User{
			Email: "test3@example.com",
			Name:  "Test User 3",
		}
		userService.On("Create", ctx, domainUser).Return(&domainUser, nil).Once()

		resp, err := server.PostUsers(ctx, req)
		require.NoError(t, err)

		//email := openapi_types.Email("test3@example.com")
		//name := "Test User 3"

		//expectedResp := PostUsers201JSONResponse(User{Email: &email, Name: &name})
		assert.NotNil(t, resp)

		userService.AssertExpectations(t)
	})
}

func TestDeleteUsersUuid(t *testing.T) {
	req := DeleteUsersUuidRequestObject{
		Uuid: uuid.New(),
	}

	t.Run("testDeleteUsersUuidPanic", func(t *testing.T) {
		assert.Panics(t, func() { server.DeleteUsersUuid(ctx, req) }, "The code did not panic")
	})
}

func TestGetUsersUuid(t *testing.T) {
	req := GetUsersUuidRequestObject{
		Uuid: uuid.New(),
	}

	t.Run("testGetUsersUuidPanic", func(t *testing.T) {
		assert.Panics(t, func() { server.GetUsersUuid(ctx, req) }, "The code did not panic")
	})
}

func TestPutUsersUuid(t *testing.T) {
	req := PutUsersUuidRequestObject{
		Uuid: uuid.New(),
		// Add other fields as needed
	}

	t.Run("testPutUsersUuidPanic", func(t *testing.T) {
		assert.Panics(t, func() { server.PutUsersUuid(ctx, req) }, "The code did not panic")
	})
}
