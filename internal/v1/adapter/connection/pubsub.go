package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog"
)

type TopicConfig struct {
	ProjectId string
	TopicId   string
}

func NewPubsubTopic(ctx context.Context, envs *config.Environment, logger *zerolog.Logger) *pubsub.Topic {
	client := NewPubsubClient(ctx, envs, logger)
	return getTopic(ctx, client, envs, logger)
}

func NewPubsubClient(ctx context.Context, envs *config.Environment, logger *zerolog.Logger) *pubsub.Client {
	if value := os.Getenv("PUBSUB_EMULATOR_HOST"); value != "" {
		logger.Info().Str("address", value).Msg("Using Pubsub Emulator")
	}
	client, err := pubsub.NewClient(ctx, envs.PubSub.ProjectId)
	if err != nil {
		logger.Fatal().Err(err).Msg(fmt.Sprintf("failed to connect to pubsub client: %v", err))
	}

	return client
}

func getTopic(ctx context.Context, client *pubsub.Client, envs *config.Environment, logger *zerolog.Logger) *pubsub.Topic {
	topic := client.Topic(envs.PubSub.TopicId)
	if x, err := topic.Exists(ctx); err != nil {
		logger.Fatal().Err(err).Msg(fmt.Sprintf("failed on topic.Exists: %v", err))
	} else if !x {
		logger.Fatal().Err(err).Msg(fmt.Sprintf("pubsub: topic %s does not exist on project %s", envs.PubSub.TopicId, envs.PubSub.ProjectId))
	}

	return topic
}
