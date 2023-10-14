package event

import (
	"github.com/google/uuid"
	"github.com/j1cs/api-user/internal/v1/domain"
)

type User struct {
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

func (u *User) FromDomain(user *domain.User) {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Email = user.Email
}
