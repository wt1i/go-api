package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const MaxPageSize = 200

// BaseModel base model
type BaseModel struct {
	ID        uint       `gorm:"id primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"deleted_at"`
}

type Pagination struct {
	Is       bool
	Page     int `form:"page"`
	PageSize int `form:"page_size" binding:"required_with=Page"`
	Offset   int
	Limit    int
}

// TODO: refactor, merge Page, Offset usage
func (p *Pagination) PagedDB(tx *gorm.DB) *gorm.DB {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = MaxPageSize
	}
	if p.PageSize == 0 {
		p.PageSize = MaxPageSize
	}

	if p.Offset == 0 {
		p.Offset = (p.Page - 1) * p.PageSize
		p.Limit = p.PageSize
	}

	return tx.Offset(p.Offset).Limit(p.Limit)
}
