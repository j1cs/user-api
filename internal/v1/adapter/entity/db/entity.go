package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Entity provides the interface that should be implemented by each entity for the database.
type Entity interface {
	GetTableName() string
	Seed(db *gorm.DB) error
}

// BaseUUID can be embedded into an entity to provide UUID, CreatedAt, and UpdatedAt fields automatically.
type BaseUUID struct {
	UUID      uuid.UUID `gorm:"<-:create;type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
}

// BeforeCreate is a hook that runs before the row is created in the database and adds a new random UUID. This should
// not be called manually as it will be called by gorm automatically.
func (m *BaseUUID) BeforeCreate(*gorm.DB) (err error) {
	m.UUID = uuid.New()
	return
}

// BaseID can be embedded into an entity to provide an auto incrementing ID, CreatedAt, and UpdatedAt fields
// automatically.
type BaseID struct {
	ID        int64     `gorm:"<-:create;primary_key;"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
}
