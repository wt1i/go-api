package repository

import (
	"context"

	"go-api/internal/domain/model"
)

// NewsRepository represent repository of  the news
// Expect implementation by the infras layer

//go:generate mockgen -source=news.go -destination=mock_news.go -package=repository
type NewsRepository interface {
	// Get obtain news by id
	Get(ctx context.Context, id uint) (*model.News, error)
	// GetAllByStatus obtain news by status
	GetAllByStatus(ctx context.Context, status model.NewsStatus, pagination model.Pagination) ([]model.News, error)
	// Save news save
	Save(ctx context.Context, news *model.News) error
	// Remove news remove by id
	Remove(ctx context.Context, id uint) error
	// Update news update by entity
	Update(ctx context.Context, news *model.News) error
}
