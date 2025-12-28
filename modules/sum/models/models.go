package models

import (
	"e-backend/internal/crud"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	// URL        string `gorm:"uniqueIndex"`
	Type string `gorm:"index"`
	Word *string
	// ParentID   *uint
	// ParentURL  *string
	// Vocabulary string `gorm:"default:sum.in.ua;not null"`
	Desc  *string
	Title *string `gorm:"index"`
}

// TableName overrides the table name used by Articles to `sum_articles`
func (Article) TableName() string {
	return "sum_articles"
}

type ArticleListFilter struct {
	crud.ListFilter
	Search string
}
