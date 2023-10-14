//go:build integration

package connection

import (
	"context"
	"os"
	"testing"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/j1cs/api-user/internal/v1/service/container"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var (
	testContainerPubsub *container.TestPubsub
)

func TestPubSub(t *testing.T) {
	testContainerPubsub = container.NewTestPubsub(t)
	ctx := context.Background()
	t.Setenv("GCP_PROJECT_ID", "test-project")
	t.Setenv("GCP_TOPIC_ID", "test-topic")
	t.Setenv("PUBSUB_EMULATOR_HOST", testContainerPubsub.ConnectionString(t))
	t.Setenv("PRETTY_LOG_FORMAT", "true")
	defer testContainerPubsub.Close(t)
	logger := log.InitializeLogger(zerolog.InfoLevel)
	conf := config.SetUpConfig(ctx, logger, "dev")
	client := NewPubsubClient(ctx, conf.Environment, logger)
	_, err := client.CreateTopic(ctx, conf.Environment.PubSub.TopicId)
	assert.Nil(t, err)
	topic := getTopic(ctx, client, conf.Environment, logger)
	assert.NotNil(t, topic)
	assert.Equal(t, os.Getenv("GCP_TOPIC_ID"), topic.ID())
}
