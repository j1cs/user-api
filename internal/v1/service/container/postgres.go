package container

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	instance *postgres.PostgresContainer
}

func NewTestDatabase(t *testing.T) *TestDatabase {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	currentDir, _ := os.Getwd()
	projectRoot, _ := findProjectRoot(currentDir, "go.mod")
	container, err := postgres.RunContainer(ctx,
		postgres.WithInitScripts(filepath.Join(projectRoot, "scripts/db", "04-testcontainer-config.sh")),
		testcontainers.WithImage("postgres:14"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForListeningPort("5432/tcp")),
	)
	if t != nil {
		require.NoError(t, err)
	}

	return &TestDatabase{
		instance: container,
	}
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "5432")
	if t != nil {
		require.NoError(t, err)
	}
	return p.Int()
}

func (db *TestDatabase) Host(t *testing.T) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	host, err := db.instance.Host(ctx)
	if t != nil {
		require.NoError(t, err)
	}
	return host
}

func (db *TestDatabase) ConnString(t *testing.T) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	connStr, err := db.instance.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	if t != nil {
		require.NoError(t, err)
	}
	return connStr
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	result := db.instance.Terminate(ctx)
	if t != nil {
		require.NoError(t, result)
	}
}

func findProjectRoot(startDir string, identifier string) (string, error) {
	currentDir := startDir
	for {
		// Check if the identifier exists in the current directory
		if _, err := os.Stat(filepath.Join(currentDir, identifier)); err == nil {
			return currentDir, nil
		}

		// Move up one directory
		newDir := filepath.Dir(currentDir)
		if newDir == currentDir {
			// We've reached the filesystem root and haven't found the identifier
			return "", fmt.Errorf("project root not found")
		}
		currentDir = newDir
	}
}
