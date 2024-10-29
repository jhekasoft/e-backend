package cv

import (
	internalModels "e-backend/internal/models"
	"e-backend/internal/modules/cv/handler"
	"e-backend/internal/modules/cv/repository"
	"e-backend/internal/modules/cv/service"
	"path"
)

const CVBaseURL = "/"
const CVDataPath = "./internal/modules/cv/data"

type CVModule struct {
}

func (m *CVModule) Name() string {
	return "CV"
}

func (m *CVModule) Run(c *internalModels.Core) error {
	repo := repository.NewRepository(CVDataPath)
	services := service.NewService(repo, CVBaseURL)
	h := handler.NewHandler(services)

	c.Echo.GET("/cv/developer-timeline", h.GetDevTimeline)
	c.Echo.GET("/cv/cv", h.GetCV)
	c.Echo.Static("/cv/public", path.Join(CVDataPath, "public"))

	return nil
}

func NewModule() internalModels.Module {
	return &CVModule{}
}
