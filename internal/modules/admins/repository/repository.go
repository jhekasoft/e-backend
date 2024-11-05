package repository

import (
	"e-backend/internal/modules/admins/models"
	ir "e-backend/internal/repository"

	"gorm.io/gorm"
)

type Repository struct {
	ir.RepositoryGeneric[models.Admin]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{*ir.NewRepository[models.Admin](db)}
}

func (r *Repository) GetMany(filter models.AdminListFilter) (items []models.Admin, err error) {
	err = r.getListQuery(filter).Offset(filter.Offset).Limit(filter.Limit).Find(&items).Error
	return
}

func (r *Repository) GetTotal(filter models.AdminListFilter) (count int64, err error) {
	err = r.getListQuery(filter).Count(&count).Error
	return
}

func (r *Repository) getListQuery(filter models.AdminListFilter) *gorm.DB {
	tx := r.GetDB().Model(&models.Admin{}).Order("id desc")

	if len(filter.Search) > 0 {
		tx.Where("name ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.Role != nil {
		tx.Where("role = ?", filter.Role)
	}
	return tx
}
