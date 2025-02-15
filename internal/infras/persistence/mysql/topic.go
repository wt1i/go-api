package mysql

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "go-api/internal/domain/model"
)

// News represent entity of the news
type Topic struct {
	BaseModel
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content" gorm:"text"`
	Status  string `json:"status"`
}

type TopicClient struct {
	DB *gorm.DB `inject:""`
}

func NewTopicClient(db *gorm.DB) *TopicClient {
	return &TopicClient{DB: db}
}

// Get topic by id return domain.topic
func (r *TopicClient) Get(ctx context.Context, id uint) (*model.Topic, error) {
	topic := &model.Topic{}
	if err := r.DB.Preload("News").First(&topic, id).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

// GetAll topic return all domain.topic
func (r *TopicClient) GetAll(ctx context.Context) ([]model.Topic, error) {
	topics := []model.Topic{}
	if err := r.DB.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

// Save to add topic
func (r *TopicClient) Save(ctx context.Context, topic *model.Topic) error {
	if err := r.DB.Save(&topic).Error; err != nil {
		return err
	}

	return nil
}

// Remove delete topic and news
func (r *TopicClient) Remove(ctx context.Context, id uint) (err error) {
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
func (r *TopicClient) Update(ctx context.Context, topic *model.Topic) error {
	if err := r.DB.Model(&topic).UpdateColumns(model.Topic{
		Name: topic.Name,
		Slug: topic.Slug,
	}).Error; err != nil {
		return err
	}

	return nil
}
