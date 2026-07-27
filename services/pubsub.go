package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub/v2"
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
// The topic name is prefixed with the current environment before publishing.
// The function returns the message ID if the message is successfully published, otherwise it returns an error.
//
// Topics are provisioned outside of this code path (as infrastructure), so this function never
// creates them on demand: publishing to a topic that does not exist just returns the error.
func (p *PubSubService) PublishMessage(ctx context.Context, topicName string, data []byte) (string, error) {
	// Get the environment.
	env := os.Getenv("ENV")

	// Add the environment as a prefix to the topic name.
	topicName = fmt.Sprintf("%s-%s", env, topicName)

	// Select a publisher for the topic.
	publisher := p.Client.Publisher(topicName)

	// Set the message.
	msg := &pubsub.Message{
		Data: data,
	}

	// Publish a message.
	msgID, err := publisher.Publish(ctx, msg).Get(ctx)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return "", err
	}

	return msgID, nil
}
