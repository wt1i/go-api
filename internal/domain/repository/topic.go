package repository

import (
	model "go-api/internal/domain/model"
)

// TopicRepository represent repository of the topic
// Expect implementation by the infras layer

//go:generate mockgen -source=topic.go -destination=mock_topic.go -package=repository
type TopicRepository interface {
	Get(id uint) (*model.Topic, error)
	GetAll() ([]model.Topic, error)
	Save(*model.Topic) error
	Remove(id uint) error
	Update(*model.Topic) error
}
