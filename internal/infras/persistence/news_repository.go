package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	model "go-api/internal/domain/model"
	"go-api/internal/domain/repository"
)

var _ repository.NewsRepository = (*NewsRepositoryImpl)(nil)

// NewsRepositoryImpl Implements repository.NewsRepository
type NewsRepositoryImpl struct {
	DB *gorm.DB `inject:""`
}

// Get news by id return domain.news
func (r *NewsRepositoryImpl) Get(id uint) (*model.News, error) {
	news := &model.News{}
	if err := r.DB.Preload("Topic").First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}

// Save to add news
func (r *NewsRepositoryImpl) Save(news *model.News) error {
	if err := r.DB.Save(&news).Error; err != nil {
		return err
	}

	return nil
}

// Remove to delete news by id
func (r *NewsRepositoryImpl) Remove(id uint) (err error) {
	tx := r.DB.Begin()

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
func (r *NewsRepositoryImpl) Update(news *model.News) error {
	n := model.News{Title: news.Title, Slug: news.Slug,
		Content: news.Content, Status: news.Status, Topic: news.Topic}
	n.UpdatedAt = news.UpdatedAt
	if err := r.DB.Model(&news).UpdateColumns(n).Error; err != nil {
		return err
	}

	return nil
}

// GetAll News return all domain.news
func (r *NewsRepositoryImpl) GetAllByStatus(status model.NewsStatus, pagination model.Pagination) ([]model.News, error) {
	var db *gorm.DB
	if status == model.NewsStatusDelete {
		db = r.DB.Unscoped()
	} else {
		db = r.DB
	}

	news := []model.News{}
	if err := pagination.PagedDB(db).Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
