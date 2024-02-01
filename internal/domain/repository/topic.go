package repository

import (
	"context"

	model "go-api/internal/domain/model"
)

// TopicRepository represent repository of the topic
// Expect implementation by the infras layer

//go:generate mockgen -source=topic.go -destination=mock_topic.go -package=repository
type TopicRepository interface {
	Get(ctx context.Context, id uint) (*model.Topic, error)
	GetAll(ctx context.Context) ([]model.Topic, error)
	Save(ctx context.Context, topic *model.Topic) error
	Remove(ctx context.Context, id uint) error
	Update(ctx context.Context, topic *model.Topic) error
}
