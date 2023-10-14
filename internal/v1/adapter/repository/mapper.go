package repository

import (
	"github.com/j1cs/api-user/internal/v1/adapter/entity/db"
	"github.com/j1cs/api-user/internal/v1/domain"
)

func UserEntitiesToDomains(users []db.User) ([]domain.User, error) {
	var domainUsers []domain.User
	for _, user := range users {
		domainUser, err := user.ToDomain()
		if err != nil {
			return nil, err
		}
		domainUsers = append(domainUsers, *domainUser)
	}
	return domainUsers, nil
}
