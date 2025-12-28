package sum

import (
	internalModels "e-backend/internal/models"
	"e-backend/modules/sum/handler"
	"e-backend/modules/sum/repository"
	"e-backend/modules/sum/service"
)

type SumModule struct {
}

func (m *SumModule) Name() string {
	return "Sum"
}

func (m *SumModule) Run(c *internalModels.Core) error {
	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	c.Echo.GET("/sum/articles", h.GetList)
	c.Echo.GET("/sum/articles/:word", h.GetWord)

	return nil
}

func NewModule() internalModels.Module {
	return &SumModule{}
}
