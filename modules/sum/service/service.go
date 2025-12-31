package service

import (
	"e-backend/internal/crud"
	"e-backend/modules/sum/models"
	"e-backend/modules/sum/repository"
	"strings"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetManyWithTotal(filter models.ArticleListFilter) (items []models.Article, total int64, err error) {
	items, err = s.repo.GetMany(filter)
	if err != nil {
		return
	}

	total, err = s.repo.GetTotal(filter)
	return
}

func (s *Service) GetWordOrAlternatives(title string) (item *models.Article, alts []string, err error) {
	// Always slice (not nil)
	alts = make([]string, 0)

	items, err := s.repo.GetMany(models.ArticleListFilter{
		ListFilter: crud.ListFilter{Limit: 10, Offset: 0},
		Search:     title,
	})
	if err != nil {
		return
	}

	if len(items) < 1 {
		return
	}

	// Exact word
	if items[0].Title != nil && strings.EqualFold(*items[0].Title, title) {
		item = &items[0]
	}

	// Alternatives
	for _, it := range items {
		if it.Word == nil {
			continue
		}
		alts = append(alts, *it.Word)
	}

	return
}
