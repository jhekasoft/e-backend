package service

import (
	"e-backend/internal/models"
)

type CRUDService[M models.CRUDModel, F any] struct {
	repo models.CRUDRepository[M, F]
}

func NewService[M models.CRUDModel, F any](repo models.CRUDRepository[M, F]) *CRUDService[M, F] {
	return &CRUDService[M, F]{repo}
}

func (s *CRUDService[M, F]) GetRepo() models.CRUDRepository[M, F] {
	return s.repo
}

func (s *CRUDService[M, F]) Create(item M) (*M, error) {
	return s.repo.Create(item)
}

func (s *CRUDService[M, F]) Update(id uint, item M) (*M, error) {
	return s.repo.Update(id, item)
}

func (s *CRUDService[M, F]) Get(id uint) (*M, error) {
	return s.repo.Get(id)
}

func (s *CRUDService[M, F]) GetManyWithTotal(filter F) (items []M, total int64, err error) {
	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *CRUDService[M, F]) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}
