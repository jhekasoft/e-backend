package service

import (
	"e-backend/internal/modules/auth/models"
	"e-backend/internal/modules/auth/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(item models.User) (*models.User, error) {
	return s.repo.Create(item)
}

func (s *Service) Update(id uint, item models.User) (*models.User, error) {
	return s.repo.Update(id, item)
}

func (s *Service) Get(id uint) (*models.User, error) {
	return s.repo.Get(id)
}

func (s *Service) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}
