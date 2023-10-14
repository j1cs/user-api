package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/port"
	"github.com/j1cs/api-user/internal/v1/service"
	"github.com/j1cs/api-user/internal/v1/service/config"
	"github.com/j1cs/api-user/internal/v1/service/log"
	"os"
)

func main() {
	ctx := context.Background()
	logger := log.InitializeLogger(zerolog.InfoLevel)
	logger.Info().Msg("Logger configured")
	conf := config.SetUpConfig(ctx, logger, os.Getenv("ENVIRONMENT"))
	logger.Info().Msg("configuration done")
	application := service.NewApplication(conf, logger)
	server := port.NewServer(application, logger)
	if err := service.RunServer(server, logger); err != nil {
		logger.Fatal().Err(err)
	}
}
