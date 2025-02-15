package adapt

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-api/internal/domain/model"
	"go-api/internal/domain/repository"
	"go-api/internal/infras/persistence/mysql"
)

var _ repository.TopicRepository = (*TopicRepositoryImpl)(nil)

// TopicRepositoryImpl Implements repository.TopicRepository
type TopicRepositoryImpl struct {
	DB *gorm.DB `inject:""`
	// Redis *redis.Client `inject:""`
}

func (t TopicRepositoryImpl) Get(ctx context.Context, id uint) (*model.Topic, error) {
	return mysql.NewTopicClient(t.DB).Get(ctx, id)
}

func (t TopicRepositoryImpl) GetAll(ctx context.Context) ([]model.Topic, error) {
	//TODO implement me
	panic("implement me")
}

func (t TopicRepositoryImpl) Save(ctx context.Context, topic *model.Topic) error {
	//TODO implement me
	panic("implement me")
}

func (t TopicRepositoryImpl) Remove(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (t TopicRepositoryImpl) Update(ctx context.Context, topic *model.Topic) error {
	//TODO implement me
	panic("implement me")
}
