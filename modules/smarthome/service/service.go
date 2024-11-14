package service

import (
	"e-backend/modules/smarthome/models"
	"e-backend/modules/smarthome/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(item models.SmartHomeSensorValue) (*models.SmartHomeSensorValue, error) {
	return s.repo.Create(item)
}

func (s *Service) Get(id uint) (*models.SmartHomeSensorValue, error) {
	return s.repo.Get(id)
}
