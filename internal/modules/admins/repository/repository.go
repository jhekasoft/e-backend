package repository

import (
	"e-backend/internal/modules/admins/models"
	ir "e-backend/internal/repository"

	"gorm.io/gorm"
)

type Repository struct {
	ir.CRUDRepository[models.Admin, models.AdminListFilter]
}

func NewRepository(db *gorm.DB) *Repository {
	listOrder := "id desc"
	var listScope ir.ListScopeFunc[models.AdminListFilter] = listScope
	return &Repository{*ir.NewRepository[models.Admin](db, &listScope, listOrder)}
}

func listScope(filter models.AdminListFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(filter.Search) > 0 {
			db.Where("name ILIKE ?", "%"+filter.Search+"%")
		}
		if filter.Role != nil {
			db.Where("role = ?", filter.Role)
		}
		return db
	}
}
