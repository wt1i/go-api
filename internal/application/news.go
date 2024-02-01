package application

import (
	"context"

	model "go-api/internal/domain/model"
	"go-api/internal/domain/repository"

	"gorm.io/gorm"
)

// NewsService news service
type NewsService struct {
	DB       *gorm.DB                  `inject:""`
	NewsRepo repository.NewsRepository `inject:""`
}

type GetNewsReq struct {
	NewsID uint `uri:"news_id"`
}

// GetNews returns domain.news by id
func (s *NewsService) GetNews(ctx context.Context, req GetNewsReq) (*model.News, error) {
	return s.NewsRepo.Get(ctx, req.NewsID)
}

type UpsertNewsReq struct {
	ID      uint             `uri:"id"`
	TopicID uint             `json:"topic_id"`
	Title   string           `json:"title"`
	Slug    string           `json:"slug"`
	Content string           `json:"content"`
	Status  model.NewsStatus `json:"status"`
}

func (a UpsertNewsReq) BuildModelNews() model.News {
	return model.News{
		TopicID: a.TopicID,
		Title:   a.Title,
		Slug:    a.Slug,
		Content: a.Content,
		Status:  a.Status,
	}
}

// AddNews saves new news
func (s *NewsService) AddNews(ctx context.Context, p model.News) error {
	return s.NewsRepo.Save(ctx, &p)
}

type RemoveNewsReq struct {
	NewsID uint `uri:"news_id"`
}

// RemoveNews do remove news by id
func (s *NewsService) RemoveNews(ctx context.Context, id uint) error {
	return s.NewsRepo.Remove(ctx, id)
}

// UpdateNews do remove news by id
func (s *NewsService) UpdateNews(ctx context.Context, id uint, p model.News) error {
	p.ID = id

	return s.NewsRepo.Update(ctx, &p)
}

type GetAllNewsReq struct {
	Status     model.NewsStatus `form:"status"`
	Pagination model.Pagination
}

// GetAllNewsByFilter return all model.News by filter
func (s *NewsService) GetAllNewsByFilter(ctx context.Context, req GetAllNewsReq) ([]model.News, error) {
	return s.NewsRepo.GetAllByStatus(ctx, req.Status, req.Pagination)
}
