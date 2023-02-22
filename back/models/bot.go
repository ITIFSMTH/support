package models

import "github.com/jinzhu/gorm"

type Bot struct {
	gorm.Model
	TGName string `gorm:"not null"`
	TGAPI  string `gorm:"unique;not null"`
}
