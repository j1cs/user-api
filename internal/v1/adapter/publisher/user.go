package publisher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/adapter/entity/event"
	"github.com/j1cs/api-user/internal/v1/domain"
	"github.com/j1cs/api-user/internal/v1/service/util"
)

type Publisher struct {
	topic  *pubsub.Topic
	logger *zerolog.Logger
}

func NewPublisher(topic *pubsub.Topic, logger *zerolog.Logger) *Publisher {
	return &Publisher{
		topic:  topic,
		logger: logger,
	}
}

func (p *Publisher) Publish(ctx context.Context, user domain.User, header domain.Header) (string, error) {
	p.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering publish adapter")
	ev := &event.User{}
	ev.FromDomain(&user)
	return p.publishMessage(ctx, ev, header)
}

func (p *Publisher) publishMessage(ctx context.Context, data interface{}, header domain.Header) (string, error) {
	rawBody, err := json.Marshal(data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("publish message: json.Marshal err %v", err))
	}

	attributes, err := util.StructToStringMap(header)
	if err != nil {
		return "", errors.New(fmt.Sprintf("publish message: attributes marshal err %v", err))
	}
	m := &pubsub.Message{
		Data:       rawBody,
		Attributes: attributes,
	}

	// whe should print the entire json: {data:{...}, attributes:{...}} maybe using logger's RawJSON
	attrJson, _ := json.Marshal(m.Attributes)
	p.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg(fmt.Sprintf("{\"data\": %s, \"attributes\": %s}", string(m.Data), attrJson))

	result := p.topic.Publish(ctx, m)

	id, err := result.Get(ctx)
	if err != nil {
		return "", errors.New(fmt.Sprintf("publish message: result.Get: %v", err))
	}

	p.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Message published: " + id + " topic " + p.topic.ID())
	return id, nil
}
