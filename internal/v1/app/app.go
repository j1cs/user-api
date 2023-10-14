package app

import "github.com/j1cs/api-user/internal/v1/domain"

type Application struct {
	Service *Service
}

type Service struct {
	User domain.UserService
}
