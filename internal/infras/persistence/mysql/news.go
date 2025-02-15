package mysql

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	model "go-api/internal/domain/model"
)

// News represent entity of the news
type News struct {
	BaseModel
	TopicID uint   `json:"topic_id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content" gorm:"text"`
	Status  string `json:"status"`
}

type NewsClient struct {
	DB *gorm.DB `inject:""`
}

func NewNewsClient(db *gorm.DB) *NewsClient {
	return &NewsClient{DB: db}
}

// Get news by id return domain.news
func (n NewsClient) Get(ctx context.Context, id uint) (*model.News, error) {
	news := &model.News{}
	if err := n.DB.Preload("Topic").First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}

// Save to add news
func (n NewsClient) Save(ctx context.Context, news *model.News) error {
	return n.DB.Save(&news).Error
}

// Remove to delete news by id
func (n NewsClient) Remove(ctx context.Context, id uint) (err error) {
	tx := n.DB.Begin()

	defer func() {
		err = TxErrDefer(tx, err)
	}()

	news := model.News{}
	if err := tx.First(&news, id).Error; err != nil {
		return err
	}

	if err := tx.Delete(&news).Error; err != nil {
		return err
	}

	return nil
}

// Update is update news
func (n NewsClient) Update(ctx context.Context, news *model.News) error {
	updates := News{
		Title:   news.Title,
		Slug:    news.Slug,
		Content: news.Content,
		Status:  string(news.Status),
	}
	updates.UpdatedAt = news.UpdatedAt
	if err := n.DB.Model(&updates).UpdateColumns(updates).Error; err != nil {
		return err
	}

	return nil
}

// GetAllByStatus GetAll News return all domain.news
func (n NewsClient) GetAllByStatus(ctx context.Context, status model.NewsStatus, pagination model.Pagination) ([]model.News, error) {
	var db *gorm.DB
	if status == model.NewsStatusDelete {
		db = n.DB.Unscoped()
	} else {
		db = n.DB
	}

	var news []model.News
	if err := pagination.PagedDB(db).Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
