package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	ChatID int64  `gorm:"not null"`
	UserID int64  `gorm:"not null"`
	Text   string `gorm:"not null"`
	Readed bool   `gorm:"not null;default:true"`
}
