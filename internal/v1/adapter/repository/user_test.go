//go:build integration

package repository

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/j1cs/api-user/internal/v1/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userRepository domain.UserRepository
	ctx            context.Context
	reqid          uint64
)

func TestUserRepository(t *testing.T) {
	userRepository = NewUserRepository(connDB, testLogger)
	requestID := GenerateRequestID("testPrefix")
	ctx = context.WithValue(context.Background(), middleware.RequestIDKey, requestID)

	t.Run("testFindAllOk", testFindAllOk)
	t.Run("testCreateOk", testCreateOk)
}

func testFindAllOk(t *testing.T) {
	users, err := userRepository.FindAll(ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, users)
	assert.Len(t, users, 3)
}

func testCreateOk(t *testing.T) {
	user := domain.User{
		Email: "jcuzmar@protonmail.cl",
		Name:  "Juan Cuzmar",
	}
	result, err := userRepository.Save(ctx, user)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, user.Email, result.Email)
}

func GenerateRequestID(prefix string) string {
	myid := atomic.AddUint64(&reqid, 1)
	return fmt.Sprintf("%s-%06d", prefix, myid)
}
