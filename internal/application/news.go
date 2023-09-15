package application

import (
	"gorm.io/gorm"

	model "github.com/liwentao0503/go-api/internal/domain/model"
	"github.com/liwentao0503/go-api/internal/domain/repository"
)

// NewsService news service
type NewsService struct {
	DB       *gorm.DB                  `inject:""`
	NewsRepo repository.NewsRepository `inject:""`
}

type GetNewsReq struct {
	NewsID uint `form:"news_id"`
}

// GetNews returns domain.news by id
func (s *NewsService) GetNews(req GetNewsReq) (*model.News, error) {
	return s.NewsRepo.Get(req.NewsID)
}

type UpsertNewsReq struct {
	ID      uint             `json:"id"`
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
func (s *NewsService) AddNews(p model.News) error {
	return s.NewsRepo.Save(&p)
}

type RemoveNewsReq struct {
	NewsID uint `form:"news_id"`
}

// RemoveNews do remove news by id
func (s *NewsService) RemoveNews(id uint) error {
	return s.NewsRepo.Remove(id)
}

// UpdateNews do remove news by id
func (s *NewsService) UpdateNews(p model.News, id uint) error {
	p.ID = id

	return s.NewsRepo.Update(&p)
}

type GetAllNewsReq struct {
	Status     model.NewsStatus `form:"status"`
	Pagination model.Pagination
}

// GetAllNewsByFilter return all model.News by filter
func (s *NewsService) GetAllNewsByFilter(req GetAllNewsReq) ([]model.News, error) {
	return s.NewsRepo.GetAllByStatus(req.Status, req.Pagination)
}
