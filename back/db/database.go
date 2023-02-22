package db

import (
	"fmt"
	"time"

	"support-back/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func InitDb(path string) {
	db, err = gorm.Open("postgres", path)

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Setting{})
	db.AutoMigrate(&models.Bot{})
	db.AutoMigrate(&models.Worker{})
	db.AutoMigrate(&models.WorkerStat{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Chat{})
	db.AutoMigrate(&models.Message{})

	created := db.First(&models.Setting{})
	if created.RowsAffected == 0 {
		db.Model(&models.Setting{}).Create(&models.Setting{
			Greeting: "Добро пожаловать! Напишите сообщение, и Вам помогут.",
		})

		db.Model(&models.Worker{}).Create(&models.Worker{
			Login:    "admin",
			Password: "$2a$14$orvjOpUxaUDzOd4Au6tBzOx1h5wbLY0/w8NE4VbSSTeScHrIEdloK",
			Role:     "admin",
		})

		db.Model(&models.WorkerStat{}).Create(&models.WorkerStat{
			Worker:   1,
			LastDate: time.Now(),
		})
	}
}

func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
