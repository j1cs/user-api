package config

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/j1cs/api-user/internal/v1/service/secretmanager"

	"github.com/rs/zerolog"
)

type SecretManager interface {
	GetSecret(ctx context.Context, secretName string) string
}

type Config struct {
	Environment *Environment
}

type Environment struct {
	ReadDatabase  *Database
	WriteDatabase *Database
	PubSub        *PubSub
}

type Database struct {
	Host       string
	Port       string
	Name       string
	SchemaName string
	User       string
	Password   string
}

type PubSub struct {
	ProjectId string
	TopicId   string
}

func SetUpConfig(ctx context.Context, logger *zerolog.Logger, env string) *Config {
	if env != "prod" {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}
	logger.Info().Msg("Setup Begin, getting environment variables")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	if env == "prod" {
		s := secretmanager.NewSecretManagerClient(ctx)
		logger.Info().Msg("getting user and pass from sm")
		user = s.GetSecret(ctx, "sql-user-capability")
		password = s.GetSecret(ctx, "sql-pass-capability")
	}
	schema := os.Getenv("SCHEMA_NAME")

	rdb := &Database{
		Host:       os.Getenv("READ_DATABASE_HOST"),
		Port:       port,
		Name:       name,
		User:       user,
		Password:   url.QueryEscape(password),
		SchemaName: schema,
	}
	wdb := &Database{
		Host:       os.Getenv("WRITE_DATABASE_HOST"),
		Port:       port,
		Name:       name,
		User:       user,
		Password:   url.QueryEscape(password),
		SchemaName: schema,
	}
	return &Config{
		Environment: &Environment{
			PubSub:        setUpPubSubEnv(),
			ReadDatabase:  rdb,
			WriteDatabase: wdb,
		},
	}
}

func setUpPubSubEnv() *PubSub {
	return &PubSub{
		ProjectId: os.Getenv("GCP_PROJECT_ID"),
		TopicId:   os.Getenv("GCP_TOPIC_ID"),
	}
}
