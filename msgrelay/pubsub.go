package msgrelay

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type PubSub interface {
	Topic(topic string) *pubsub.Topic
	Publish(ctx context.Context,
		topic *pubsub.Topic,
		data []byte,
	) *pubsub.PublishResult
}

type PubSubClient struct {
	Client *pubsub.Client
}

func NewPubSub(ctx context.Context, projectID string) (*PubSubClient, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &PubSubClient{Client: client}, err
}

func (p *PubSubClient) Topic(topic string) *pubsub.Topic {
	return p.Client.Topic(topic)
}

func (p *PubSubClient) Publish(
	ctx context.Context,
	topic *pubsub.Topic,
	data []byte,
) *pubsub.PublishResult {
	return topic.Publish(ctx, &pubsub.Message{Data: data})
}
