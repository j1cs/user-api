package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestEntity checks that the base entity struct and functions work as expected.
func TestEntity(t *testing.T) {
	t.Run("UUID generation", testEntityUUIDGeneration)
}

// testEntityUUIDGeneration will make sure that new entities have a uniquely generated UUID.
func testEntityUUIDGeneration(t *testing.T) {
	var entity = BaseUUID{}
	err := entity.BeforeCreate(&gorm.DB{})
	assert.Nil(t, err)
	assert.NotEqual(t, uuid.Nil, entity)
}
