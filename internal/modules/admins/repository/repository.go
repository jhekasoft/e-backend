package repository

import (
	"e-backend/internal/modules/admins/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(item models.Admin) (createdItem *models.Admin, err error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	createdItem, err = r.Get(item.ID)
	return
}

func (r *Repository) Update(id uint, item models.Admin) (*models.Admin, error) {
	var updatedItem models.Admin
	if err := r.db.Where("id = ?", id).Updates(&item).Scan(&updatedItem).Error; err != nil {
		return nil, err
	}

	return &updatedItem, nil
}

func (r *Repository) Get(id uint) (item *models.Admin, err error) {
	err = r.db.First(&item, id).Error
	return
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
	tx := r.db.Model(&models.Admin{}).Order("id desc")

	if len(filter.Search) > 0 {
		tx.Where("name ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.Role != nil {
		tx.Where("role = ?", filter.Role)
	}
	return tx
}

func (r *Repository) Delete(id uint) (err error) {
	err = r.db.Delete(&models.Admin{}, id).Error
	return
}
