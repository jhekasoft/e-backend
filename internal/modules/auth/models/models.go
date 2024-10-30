package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string     `gorm:"uniqueIndex"`
	Email       string     `gorm:"uniqueIndex"`
	Name        string     `gorm:"index"`
	ActivatedAt *time.Time `gorm:"index"`
	Password    string     `json:"-"`
}
