//go:build integration

package connection

import (
	"context"
	"fmt"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"

	"github.com/j1cs/api-user/internal/v1/service/container"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var testDB *container.TestDatabase

func singleContainerConnect(t *testing.T) {
	if testDB == nil {
		testDB = container.NewTestDatabase(t)
		time.Sleep(time.Second * 3)
	}
}

func TestDatabase(t *testing.T) {
	singleContainerConnect(t)
	t.Setenv("WRITE_DATABASE_HOST", testDB.Host(t))
	t.Setenv("READ_DATABASE_HOST", testDB.Host(t))
	t.Setenv("DATABASE_PORT", fmt.Sprint(testDB.Port(t)))
	t.Setenv("DATABASE_NAME", "postgres")
	t.Setenv("DATABASE_USER", "postgres")
	t.Setenv("DATABASE_PASSWORD", "postgres")
	t.Setenv("PRETTY_LOG_FORMAT", "true")

	defer testDB.Close(t)

	ctx := context.Background()
	logger := log.InitializeLogger(zerolog.InfoLevel)
	conf := config.SetUpConfig(ctx, logger, "dev")
	rw := NewDatabaseConnection(conf.Environment, logger)
	assert.NotNil(t, rw)
	assert.Nil(t, rw.Error)

	ro := NewDatabaseConnection(conf.Environment, logger)
	assert.NotNil(t, ro)
	assert.Nil(t, ro.Error)
}
