package models

import "gorm.io/gorm"

type Model interface {
	GetID() uint
}

type Repository[M Model, F any] interface {
	GetDB() *gorm.DB
	Create(item M) (createdItem *M, err error)
	Update(id uint, item M) (*M, error)
	Get(id uint) (item *M, err error)
	GetMany(filter F) (items []M, err error)
	GetTotal(filter F) (count int64, err error)
	Delete(id uint) (err error)
}

type Service[M Model, F any] interface {
	GetRepo() Repository[M, F]
	Create(item M) (createdItem *M, err error)
	Update(id uint, item M) (*M, error)
	Get(id uint) (item *M, err error)
	GetManyWithTotal(filter F) (items []M, total int64, err error)
	Delete(id uint) (err error)
}
