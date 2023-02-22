package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Chat struct {
	gorm.Model
	UserID   int64         `gorm:"unique, not null"`
	Messages pq.Int64Array `gorm:"type:int[];default:'{}'"`
	Users    pq.Int64Array `gorm:"type:bigint[];default:'{}'"`
}
