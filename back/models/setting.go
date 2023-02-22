package models

import (
	"github.com/jinzhu/gorm"
)

type Setting struct {
	gorm.Model
	Greeting string `gorm:"unique, not null"`
}
