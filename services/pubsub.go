package services

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

type PubSubService struct {
	Client *pubsub.Client
}

func NewPubSubService(client *pubsub.Client) *PubSubService {
	return &PubSubService{
		Client: client,
	}
}

func (p *PubSubService) PublishMessage(topicName string, data []byte) (string, error) {
	ctx := context.Background()

	// Select a topic.
	topic := p.Client.Topic(topicName)

	// Set the message.
	msg := &pubsub.Message{
		Data: data,
	}

	// Publish a message.
	msgID, err := topic.Publish(ctx, msg).Get(ctx)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return "", err
	}

	return msgID, nil
}

func (p *PubSubService) Subscribe(topicName string, subName string, handler func(context.Context, *pubsub.Message)) error {
	ctx := context.Background()

	// Select a topic.
	topic := p.Client.Topic(topicName)

	// Create a new subscription.
	sub, err := p.Client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		log.Printf("Failed to create subscription: %v", err)
		return err
	}

	// Receive messages for subscription.
	err = sub.Receive(ctx, handler)
	if err != nil {
		log.Printf("Failed to receive message: %v", err)
		return err
	}

	return nil
}
