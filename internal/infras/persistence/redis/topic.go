package redis

import (
	"context"
	"github.com/go-redis/redis/v8"

	"go-api/internal/domain/model"
)

type TopicClient struct {
	redis *redis.Client
}

func NewTopicClient(redisClient *redis.Client) *TopicClient {
	return &TopicClient{redis: redisClient}
}

// Get topic by id return domain.topic
func (t *TopicClient) Get(ctx context.Context, id uint) (*model.Topic, error) {
	return nil, nil
}

// GetAll topic return all domain.topic
func (t *TopicClient) GetAll(ctx context.Context) ([]model.Topic, error) {
	return nil, nil
}

// Save to add topic
func (t *TopicClient) Save(ctx context.Context, topic *model.Topic) error {
	return nil
}

// Remove delete topic and news
func (t *TopicClient) Remove(ctx context.Context, id uint) (err error) {
	return nil
}

// Update data topic
func (t *TopicClient) Update(ctx context.Context, topic *model.Topic) error {
	return nil
}
