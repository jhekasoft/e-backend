package auth

import (
	internalModels "e-backend/internal/models"
	"e-backend/internal/modules/auth/handler"
	"e-backend/internal/modules/auth/models"
	"e-backend/internal/modules/auth/repository"
	"e-backend/internal/modules/auth/service"
)

type AuthModule struct {
}

func (m *AuthModule) Name() string {
	return "Auth"
}

func (m *AuthModule) Run(c *internalModels.Core) error {
	c.DB.AutoMigrate(&models.User{})

	repo := repository.NewRepository(c.DB)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	c.Echo.POST("/auth/user", h.CreateItem)

	return nil
}

func NewModule() internalModels.Module {
	return &AuthModule{}
}
