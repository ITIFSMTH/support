package models

import (
	"github.com/jinzhu/gorm"
)

type Worker struct {
	gorm.Model
	Login    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null;default:'operator'"`
}
