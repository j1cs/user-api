package db

import (
	"sync"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/j1cs/api-user/internal/v1/domain"
)

// User is for identifying a person in the system.
type User struct {
	BaseUUID
	Name  string
	Email string
}

// GetTableName returns the name of the table.
func (*User) GetTableName() string {
	s, _ := schema.Parse(&User{}, &sync.Map{}, config.GormConf().NamingStrategy)
	return s.Table
}

// Seed will fill the database with example data for development.
func (*User) Seed(db *gorm.DB) error {
	users := []User{
		{
			Name:  "Alice",
			Email: "alice@example.com",
		},
		{
			Name:  "Bob",
			Email: "bob@example.com",
		},
		{
			Name:  "Charlie",
			Email: "charlie@example.com",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

// ToDomain converts the database entity into a domain object.
func (u *User) ToDomain() (*domain.User, error) {
	user, err := domain.NewIdentifiedUser(u.UUID, u.Email, u.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FromDomain converts a user domain object into a database entity.
func (u *User) FromDomain(user *domain.User) {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Email = user.Email
}
