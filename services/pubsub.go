package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

// PubSubService represents a service for interacting with a Pub/Sub system.
type PubSubService struct {
	Client *pubsub.Client
}

// NewPubSubService creates a new instance of the PubSubService using the provided pubsub.Client.
func NewPubSubService(client *pubsub.Client) *PubSubService {
	return &PubSubService{
		Client: client,
	}
}

// PublishMessage publishes a message to a specified topic in the Pub/Sub service.
// It takes a topicName string and data []byte as parameters and returns the message ID and an error (if any).
// If the topic doesn't exist, it creates a new topic with the given name.
// The function first checks if the topic exists, and if not, it creates the topic.
// Then, it sets the message data and publishes the message to the topic.
// The function returns the message ID if the message is successfully published, otherwise it returns an error.
func (p *PubSubService) PublishMessage(topicName string, data []byte) (string, error) {
	ctx := context.Background()

	// Get the environment.
	env := os.Getenv("ENV")

	// Add the environment as a prefix to the topic name.
	topicName = fmt.Sprintf("%s-%s", env, topicName)

	// Select a topic.
	topic := p.Client.Topic(topicName)

	// Check if the topic exists.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Printf("Failed to check if topic exists: %v", err)
		return "", err
	}

	// If the topic doesn't exist, create it.
	if !exists {
		topic, err = p.Client.CreateTopic(ctx, topicName)
		if err != nil {
			log.Printf("Failed to create topic: %v", err)
			return "", err
		}
	}

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

// Subscribe creates a new subscription for a given topic and starts receiving messages.
// It takes the topic name, subscription name, and a message handler function as parameters.
// The message handler function is responsible for processing the received messages.
// It returns an error if there was a problem creating the subscription or receiving messages.
func (p *PubSubService) Subscribe(topicName string, subName string, handler func(context.Context, *pubsub.Message)) error {
	ctx := context.Background()

	// Get the environment.
	env := os.Getenv("ENV")

	// Add the environment as a prefix to the topic name and subscription name.
	topicName = fmt.Sprintf("%s-%s", env, topicName)
	subName = fmt.Sprintf("%s-%s", env, subName)

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
