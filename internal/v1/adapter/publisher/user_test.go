//go:build integration

package publisher

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/j1cs/api-user/internal/v1/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userPublisher domain.UserPublisher
	reqid         uint64
	currentCtx    context.Context
)

func TestUserPublisher(t *testing.T) {
	userPublisher = NewPublisher(testTopic, testLogger)
	requestID := GenerateRequestID("testPrefix")
	currentCtx = context.WithValue(context.Background(), middleware.RequestIDKey, requestID)

	t.Run("testPublishOk", testPublishOk)
}

func testPublishOk(t *testing.T) {
	user := domain.User{
		UUID:  uuid.New(),
		Email: "jcuzmar@protonmail.cl",
		Name:  "Juan Cuzmar",
	}
	header := domain.NewHeader()
	result, err := userPublisher.Publish(currentCtx, user, header)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func GenerateRequestID(prefix string) string {
	myid := atomic.AddUint64(&reqid, 1)
	return fmt.Sprintf("%s-%06d", prefix, myid)
}
