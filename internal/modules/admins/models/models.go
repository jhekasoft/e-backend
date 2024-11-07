package models

import (
	ir "e-backend/internal/repository"
)

// AdminRole is role of the administrator.
type AdminRole string

const (
	AdminRoleSuper   AdminRole = "super"
	AdminRoleDefault AdminRole = "default"
)

type Admin struct {
	ir.CRUDModel
	Username string    `gorm:"uniqueIndex"`
	Name     string    `gorm:"index"`
	Role     AdminRole `gorm:"index"`
	Password string    `json:"-"`
}

type AdminListFilter struct {
	ir.ListFilter
	Role   *AdminRole
	Search string
}
