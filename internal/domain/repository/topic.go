package repository

import (
	model "github.com/liwentao0503/go-api/internal/domain/model"
)

// TopicRepository represent repository of the topic
// Expect implementation by the infras layer
type TopicRepository interface {
	Get(id uint) (*model.Topic, error)
	GetAll() ([]model.Topic, error)
	Save(*model.Topic) error
	Remove(id uint) error
	Update(*model.Topic) error
}
