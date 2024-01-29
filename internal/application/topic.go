package application

import (
	"context"
	model "go-api/internal/domain/model"
	"go-api/internal/domain/repository"
)

// TopicService topic service
type TopicService struct {
	TopicRepo repository.TopicRepository `inject:""`
}

type GetTopicReq struct {
	TopicID uint `uri:"topic_id"`
}

// GetTopic returns a topic by id
func (s *TopicService) GetTopic(ctx context.Context, id uint) (*model.Topic, error) {
	return s.TopicRepo.Get(ctx, id)
}

// GetAllTopic return all topics
func (s *TopicService) GetAllTopic(ctx context.Context) ([]model.Topic, error) {
	return s.TopicRepo.GetAll(ctx)
}

type UpsertTopicReq struct {
	ID   uint   `uri:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (a UpsertTopicReq) BuildModelTopic() model.Topic {
	return model.Topic{
		Name: a.Name,
		Slug: a.Slug,
	}
}

// AddTopic saves new topic
func (s *TopicService) AddTopic(ctx context.Context, name string, slug string) error {
	u := &model.Topic{
		Name: name,
		Slug: slug,
	}
	return s.TopicRepo.Save(ctx, u)
}

type RemoveTopicReq struct {
	TopicID uint `uri:"topic_id"`
}

// RemoveTopic do remove topic by id
func (s *TopicService) RemoveTopic(ctx context.Context, id uint) error {
	return s.TopicRepo.Remove(ctx, id)
}

// UpdateTopic do update topic by id
func (s *TopicService) UpdateTopic(ctx context.Context, id uint, topic model.Topic) error {
	topic.ID = id
	return s.TopicRepo.Update(ctx, &topic)
}
