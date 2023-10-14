package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestUser will make sure that the user entity behaves as expected.
func TestUser(t *testing.T) {
	t.Run("Table name", testUserTableName)
	t.Run("UUID generation", testUserUUIDGeneration)
}

// testUserTableName checks that the table name is the same one that is expected.
func testUserTableName(t *testing.T) {
	var user User
	tableName := user.GetTableName()
	assert.Equal(t, "user", tableName)
}

// testUserUUIDGeneration will check that a newly generated user has a UUID.
func testUserUUIDGeneration(t *testing.T) {
	var user = User{}
	err := user.BeforeCreate(&gorm.DB{})
	assert.Nil(t, err)
	assert.NotEqual(t, uuid.Nil, user)
}
