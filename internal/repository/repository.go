package repository

import (
	"e-backend/internal/models"

	"gorm.io/gorm"
)

type RepositoryGeneric[T models.Model] struct {
	db *gorm.DB
}

func NewRepository[T models.Model](db *gorm.DB) *RepositoryGeneric[T] {
	return &RepositoryGeneric[T]{db}
}

func (r *RepositoryGeneric[T]) GetDB() *gorm.DB {
	return r.db
}

func (r *RepositoryGeneric[T]) Create(item T) (createdItem *T, err error) {
	if err := r.db.Create(&item).Error; err != nil {
		return nil, err
	}

	createdItem, err = r.Get(item.GetID())
	return
}

func (r *RepositoryGeneric[T]) Update(id uint, item T) (*T, error) {
	var updatedItem T
	if err := r.db.Where("id = ?", id).Updates(&item).Scan(&updatedItem).Error; err != nil {
		return nil, err
	}

	return &updatedItem, nil
}

func (r *RepositoryGeneric[T]) Get(id uint) (item *T, err error) {
	err = r.db.First(&item, id).Error
	return
}

func (r *RepositoryGeneric[T]) Delete(id uint) (err error) {
	var item T
	err = r.db.Delete(&item, id).Error
	return
}
