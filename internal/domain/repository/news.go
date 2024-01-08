package repository

import (
	"go-api/internal/domain/model"
)

// NewsRepository represent repository of  the news
// Expect implementation by the infras layer

//go:generate mockgen -source=news.go -destination=../mock/mock_news.go -package=mock
type NewsRepository interface {
	// Get obtain news by id
	Get(id uint) (*model.News, error)
	// GetAllByStatus obtain news by status
	GetAllByStatus(status model.NewsStatus, pagination model.Pagination) ([]model.News, error)
	// Save news save
	Save(*model.News) error
	// Remove news remove by id
	Remove(id uint) error
	// Update news update by entity
	Update(*model.News) error
}
