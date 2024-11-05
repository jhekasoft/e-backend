package service

import (
	"e-backend/internal/models"
)

type ServiceGeneric[M models.Model, F any] struct {
	repo models.Repository[M, F]
}

func NewService[M models.Model, F any](repo models.Repository[M, F]) *ServiceGeneric[M, F] {
	return &ServiceGeneric[M, F]{repo}
}

func (s *ServiceGeneric[M, F]) GetRepo() models.Repository[M, F] {
	return s.repo
}

func (s *ServiceGeneric[M, F]) Create(item M) (*M, error) {
	return s.repo.Create(item)
}

func (s *ServiceGeneric[M, F]) Update(id uint, item M) (*M, error) {
	return s.repo.Update(id, item)
}

func (s *ServiceGeneric[M, F]) Get(id uint) (*M, error) {
	return s.repo.Get(id)
}

func (s *ServiceGeneric[M, F]) GetManyWithTotal(filter F) (items []M, total int64, err error) {
	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *ServiceGeneric[M, F]) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}
