package container

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestPubsub struct {
	instance testcontainers.Container
}

func NewTestPubsub(t *testing.T) *TestPubsub {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "gcr.io/google.com/cloudsdktool/cloud-sdk:367.0.0-emulators",
		ExposedPorts: []string{"8085/tcp"},
		Cmd: []string{
			"/bin/sh",
			"-c",
			"gcloud beta emulators pubsub start --host-port 0.0.0.0:8085",
		},
		WaitingFor: wait.ForLog("started"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if t != nil {
		require.NoError(t, err)
	}

	return &TestPubsub{instance: container}
}

func (tp *TestPubsub) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := tp.instance.MappedPort(ctx, "8085")
	if t != nil {
		require.NoError(t, err)
	}
	return p.Int()
}

func (tp *TestPubsub) Host(t *testing.T) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	host, err := tp.instance.Host(ctx)
	if t != nil {
		require.NoError(t, err)
	}
	return host
}

func (tp *TestPubsub) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("%s:%d", tp.Host(t), tp.Port(t))
}

func (tp *TestPubsub) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	result := tp.instance.Terminate(ctx)
	if t != nil {
		require.NoError(t, result)
	}
}
