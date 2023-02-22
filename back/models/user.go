package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	TGID     int64  `gorm:"unique, not null"`
	BotID    int    `gorm:"not null"`
	Username string `gorm:"not null"`
	TGLink   string `gorm:"unique"`
}
