package persistence

import (
	"context"

	model "go-api/internal/domain/model"
	"go-api/internal/domain/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _ repository.TopicRepository = (*TopicRepositoryImpl)(nil)

// TopicRepositoryImpl Implements repository.TopicRepository
type TopicRepositoryImpl struct {
	DB *gorm.DB `inject:""`
}

// Get topic by id return domain.topic
func (r *TopicRepositoryImpl) Get(ctx context.Context, id uint) (*model.Topic, error) {
	topic := &model.Topic{}
	if err := r.DB.Preload("News").First(&topic, id).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

// GetAll topic return all domain.topic
func (r *TopicRepositoryImpl) GetAll(ctx context.Context) ([]model.Topic, error) {
	topics := []model.Topic{}
	if err := r.DB.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

// Save to add topic
func (r *TopicRepositoryImpl) Save(ctx context.Context, topic *model.Topic) error {
	if err := r.DB.Save(&topic).Error; err != nil {
		return err
	}

	return nil
}

// Remove delete topic and news
func (r *TopicRepositoryImpl) Remove(ctx context.Context, id uint) (err error) {
	tx := r.DB.Begin()

	defer func() {
		err = TxErrDefer(tx, err)
	}()

	topic := &model.Topic{}
	if err := r.DB.First(&topic, id).Error; err != nil {
		return err
	}

	if err := r.DB.Where("topic_id = ?", topic.ID).Delete(&model.News{}).Error; IsDBError(err) {
		return err
	}

	if err := r.DB.Delete(&topic).Error; err != nil {
		return err
	}

	return nil
}

// Update data topic
func (r *TopicRepositoryImpl) Update(ctx context.Context, topic *model.Topic) error {
	if err := r.DB.Model(&topic).UpdateColumns(model.Topic{
		Name: topic.Name,
		Slug: topic.Slug,
	}).Error; err != nil {
		return err
	}

	return nil
}
