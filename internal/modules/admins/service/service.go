package service

import (
	im "e-backend/internal/models"
	"e-backend/internal/modules/admins/models"
	is "e-backend/internal/service"
)

type Service struct {
	is.ServiceGeneric[models.Admin, models.AdminListFilter]
}

func NewService(repo im.Repository[models.Admin, models.AdminListFilter]) *Service {
	return &Service{*is.NewService(repo)}
}
