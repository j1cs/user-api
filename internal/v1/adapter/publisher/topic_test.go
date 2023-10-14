//go:build integration

package publisher

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/j1cs/api-user/internal/v1/service/config"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/adapter/connection"
	"github.com/j1cs/api-user/internal/v1/service/container"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

var (
	testContainerPubsub *container.TestPubsub
	testLogger          *zerolog.Logger
	testClient          *pubsub.Client
	testTopic           *pubsub.Topic
)

func TestMain(m *testing.M) {
	testContainerPubsub = container.NewTestPubsub(nil)
	ctx := context.Background()
	os.Setenv("GCP_PROJECT_ID", "test-project")
	os.Setenv("GCP_TOPIC_ID", "test-topic")
	os.Setenv("PUBSUB_EMULATOR_HOST", testContainerPubsub.ConnectionString(nil))
	os.Setenv("PRETTY_LOG_FORMAT", "true")
	testLogger = log.InitializeLogger(zerolog.InfoLevel)
	conf := config.SetUpConfig(ctx, testLogger, "dev")
	testClient = connection.NewPubsubClient(ctx, conf.Environment, testLogger)
	_, err := testClient.CreateTopic(ctx, conf.Environment.PubSub.TopicId)
	if err != nil {
		panic(fmt.Sprintf("failed to create topic: %v", err))
	}
	testTopic = connection.NewPubsubTopic(ctx, conf.Environment, testLogger)
	code := m.Run()
	testContainerPubsub.Close(nil)
	os.Exit(code)
}
