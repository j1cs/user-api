package service

import (
	"context"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"github.com/j1cs/api-user/internal/v1/adapter/publisher"
	"github.com/j1cs/api-user/internal/v1/adapter/repository"
	"github.com/j1cs/api-user/internal/v1/app/service"

	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/adapter/connection"
	"github.com/j1cs/api-user/internal/v1/app"
)

func NewApplication(config *config.Config, logger *zerolog.Logger) *app.Application {
	ctx := context.Background()
	db := connection.NewDatabaseConnection(config.Environment, logger)
	topic := connection.NewPubsubTopic(ctx, config.Environment, logger)
	pub := publisher.NewPublisher(topic, logger)
	userRepository := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepository, pub, logger)
	return &app.Application{
		Service: &app.Service{User: userService},
	}
}
