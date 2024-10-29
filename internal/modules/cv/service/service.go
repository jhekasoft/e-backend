package service

import (
	"e-backend/internal/modules/cv/models"
	"e-backend/internal/modules/cv/repository"
)

type Service struct {
	repo      *repository.Repository
	cvBaseURL string
}

func NewService(repo *repository.Repository, cvBaseURL string) *Service {
	return &Service{repo, cvBaseURL}
}

func (s *Service) GetDevTimeline(reverse bool) ([]models.DevTimelineItem, error) {
	return s.repo.GetDevTimeline(reverse)
}

func (s *Service) GetCVCommon() (*models.CVCommon, error) {
	return s.repo.GetCVCommon()
}

func (s *Service) GetCVEducation() ([]models.CVEducationItem, error) {
	return s.repo.GetCVEducation()
}

func (s *Service) GetCVExperience() ([]models.CVExperienceItem, error) {
	return s.repo.GetCVExperience()
}

func (s *Service) GetCVPublications() ([]models.CVPublication, error) {
	list, err := s.repo.GetCVPublications()
	if err != nil {
		return nil, err
	}

	for key, item := range list {
		item.ImageURL = s.cvBaseURL + item.ImageURL
		list[key] = item
	}

	return list, nil
}

func (s *Service) GetCVSoftwareProjects() ([]models.CVSoftwareProject, error) {
	list, err := s.repo.GetCVSoftwareProjects()

	if err != nil {
		return nil, err
	}

	for key, item := range list {
		item.ImageURL = s.cvBaseURL + item.ImageURL
		list[key] = item
	}

	return list, nil
}
