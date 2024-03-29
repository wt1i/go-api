package migration

import (
	"log"

	model "go-api/internal/domain/model"

	"github.com/jinzhu/gorm"
)

// MigrateAction db migrate action
type MigrateAction struct {
	DB *gorm.DB `inject:""`
}

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func (m *MigrateAction) DBMigrate() error {
	err := m.DB.AutoMigrate(model.News{}, model.Topic{}).Error
	log.Println("Migration error: ", err)
	log.Println("Migration has been processed")

	return err
}
