package adapt

import (
	"context"
	"go-api/internal/domain/model"
	"go-api/internal/domain/repository"
	"go-api/internal/infras/persistence/mysql"

	"github.com/jinzhu/gorm"
)

var _ repository.NewsRepository = (*NewsRepositoryImpl)(nil)

// NewsRepositoryImpl Implements repository.NewsRepository
type NewsRepositoryImpl struct {
	DB *gorm.DB `inject:""`
	// Redis *redis.Client `inject:""`
}

func (n NewsRepositoryImpl) Get(ctx context.Context, id uint) (*model.News, error) {
	return mysql.NewNewsClient(n.DB).Get(ctx, id)
}

func (n NewsRepositoryImpl) GetAllByStatus(ctx context.Context, status model.NewsStatus, pagination model.Pagination) ([]model.News, error) {
	return mysql.NewNewsClient(n.DB).GetAllByStatus(ctx, status, pagination)
}

func (n NewsRepositoryImpl) Save(ctx context.Context, news *model.News) error {
	return mysql.NewNewsClient(n.DB).Save(ctx, news)
}

func (n NewsRepositoryImpl) Remove(ctx context.Context, id uint) error {
	return mysql.NewNewsClient(n.DB).Remove(ctx, id)
}

func (n NewsRepositoryImpl) Update(ctx context.Context, news *model.News) error {
	return mysql.NewNewsClient(n.DB).Update(ctx, news)
}
