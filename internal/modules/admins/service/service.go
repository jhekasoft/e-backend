package service

import (
	"e-backend/internal/modules/admins/models"
	"e-backend/internal/modules/admins/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(item models.Admin) (*models.Admin, error) {
	return s.repo.Create(item)
}

func (s *Service) Update(id uint, item models.Admin) (*models.Admin, error) {
	return s.repo.Update(id, item)
}

func (s *Service) Get(id uint) (*models.Admin, error) {
	return s.repo.Get(id)
}

func (s *Service) GetManyWithTotal(filter models.AdminListFilter) (items []models.Admin, total int64, err error) {
	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *Service) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}
