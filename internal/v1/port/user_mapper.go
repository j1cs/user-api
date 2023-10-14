package port

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/j1cs/api-user/internal/v1/domain"
)

func (u *UserInput) ToDomain() (domain.User, error) {
	return domain.NewUser(string(u.Email), u.Name)
}

func (u *User) FromDomain(user *domain.User) {
	id := user.UUID.String()
	email := openapi_types.Email(user.Email)
	u.Uuid = &id
	u.Name = &user.Name
	u.Email = &email
}

func UserDomainsToUsers(domainUsers []domain.User) []User {
	var users []User
	for _, domainUser := range domainUsers {
		var user = User{}
		user.FromDomain(&domainUser)
		users = append(users, user)
	}
	return users
}
