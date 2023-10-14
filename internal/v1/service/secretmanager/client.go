package secretmanager

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type SecretManagerClient struct {
	projectId string
	sm        *secretmanager.Client
}

func NewSecretManagerClient(ctx context.Context) *SecretManagerClient {
	projectId := os.Getenv("GCP_PROJECT_ID")
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup secret manager client: %v", err)
	}
	return &SecretManagerClient{projectId, client}
}

func (s *SecretManagerClient) GetSecret(ctx context.Context, secretName string) string {
	// Build the request
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", s.projectId, secretName),
	}
	// Call the API.
	result, err := s.sm.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}
	return string(result.Payload.Data)
}
