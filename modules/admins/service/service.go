package service

import (
	"e-backend/internal/crud"
	"e-backend/modules/admins/models"
)

type Service struct {
	crud.Service[models.Admin, models.AdminListFilter]
}

func NewService(repo crud.CRUDRepository[models.Admin, models.AdminListFilter]) *Service {
	return &Service{*crud.NewService(repo)}
}
