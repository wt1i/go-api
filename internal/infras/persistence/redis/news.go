package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-api/internal/domain/model"
)

type NewsClient struct {
	redisClient redis.Client
}

func NewNewsClient(redis redis.Client) *NewsClient {
	return &NewsClient{redisClient: redis}
}

// Get news by id return domain.news
func (n NewsClient) Get(ctx context.Context, id uint) (*model.News, error) {
	return nil, nil
}

// Save to add news
func (n NewsClient) Save(ctx context.Context, news *model.News) error {
	return nil
}

// Remove to delete news by id
func (n NewsClient) Remove(ctx context.Context, id uint) (err error) {
	return nil
}

// Update is update news
func (n NewsClient) Update(ctx context.Context, news *model.News) error {
	return nil
}

// GetAllByStatus GetAll News return all domain.news
func (n NewsClient) GetAllByStatus(ctx context.Context, status model.NewsStatus, pagination model.Pagination) ([]model.News, error) {
	return nil, nil
}
