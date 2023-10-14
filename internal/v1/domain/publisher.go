package domain

import "context"

type UserPublisher interface {
	Publish(ctx context.Context, user User, header Header) (string, error)
}
