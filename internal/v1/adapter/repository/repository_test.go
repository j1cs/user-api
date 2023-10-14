//go:build integration

package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"github.com/j1cs/api-user/internal/v1/adapter/connection"
	"github.com/j1cs/api-user/internal/v1/service/container"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var (
	testDB      *container.TestDatabase
	testWriteDB *gorm.DB
	connDB      *gorm.DB
	testLogger  *zerolog.Logger
)

func TestMain(m *testing.M) {
	testDB = container.NewTestDatabase(nil)
	os.Setenv("WRITE_DATABASE_HOST", testDB.Host(nil))
	os.Setenv("READ_DATABASE_HOST", testDB.Host(nil))
	os.Setenv("DATABASE_PORT", fmt.Sprint(testDB.Port(nil)))
	os.Setenv("DATABASE_NAME", "postgres")
	os.Setenv("DATABASE_USER", "postgres")
	os.Setenv("SCHEMA_NAME", "public")
	os.Setenv("DATABASE_PASSWORD", "postgres")
	os.Setenv("PRETTY_LOG_FORMAT", "true")
	testLogger = log.InitializeLogger(zerolog.InfoLevel)
	conf := config.SetUpConfig(ctx, testLogger, "dev")
	connDB = connection.NewDatabaseConnection(conf.Environment, testLogger)
	code := m.Run()

	testDB.Close(nil)

	os.Exit(code)
}
