package models

import "gorm.io/gorm"

// AdminRole is role of the administrator.
type AdminRole string

const (
	AdminRoleSuper   AdminRole = "super"
	AdminRoleDefault AdminRole = "default"
)

type Admin struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex"`
	Name     string    `gorm:"index"`
	Role     AdminRole `gorm:"index"`
	Password string    `json:"-"`
}

type AdminListFilter struct {
	Offset int
	Limit  int
	Role   *AdminRole
	Search string
}
