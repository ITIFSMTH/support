package models

import (
	"time"

	strftime "github.com/itchyny/timefmt-go"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type WorkerStat struct {
	Worker   uint          `gorm:"not null"`
	Messages pq.Int64Array `gorm:"type:int[];not null;default:'{0, 0, 0, 0, 0}'"`
	Logins   pq.Int64Array `gorm:"type:bigint[];default:'{}'"`
	LastDate time.Time     `gorm:"not null;"`
}

func (WorkerStat) NewMessage(db *gorm.DB, worker uint) {
	// Get stat
	var stat WorkerStat
	db.Model(&stat).Where("worker = ?", worker).First(&stat)

	// Check is stat today
	if newDate := time.Now(); stat.LastDate.Day() != newDate.Day() {
		daysBetween := 0
		checkDate := newDate
		for {
			checkDate = checkDate.AddDate(0, 0, -1)
			daysBetween++

			if strftime.Format(checkDate, "%Y%m%d") == strftime.Format(stat.LastDate, "%Y%m%d") {
				break
			}
		}
		newMessages := stat.Messages[daysBetween:]
		for i := 4 - len(newMessages); i != 0; i-- {
			newMessages = append(newMessages, 0)
		}
		newMessages = append(newMessages, 1)

		db.Model(&stat).Where("worker = ?", worker).Updates(&WorkerStat{
			Messages: newMessages,
			LastDate: newDate,
		})
		return
	}

	// Increment messages today
	newMessages := stat.Messages[1:]
	newMessages = append(newMessages, stat.Messages[len(stat.Messages)-1]+1)
	db.Model(&stat).Where("worker = ?", worker).Updates(&WorkerStat{Messages: newMessages})
}

func (WorkerStat) NewLogin(db *gorm.DB, worker uint) {
	// Get stat
	var stat WorkerStat
	db.Model(&stat).Where("worker = ?", worker).First(&stat)

	// Get new logins (Delete old and insert new login)
	var newLogins []int64
	for _, l := range stat.Logins {
		lt := time.UnixMilli(l)
		if time.Now().AddDate(0, 0, -5).Unix() < lt.UnixMilli() {
			newLogins = append(newLogins, l)
		}
	}
	newLogins = append(newLogins, time.Now().Unix())

	db.Model(&stat).Where("worker = ?", worker).Updates(&WorkerStat{
		Logins: newLogins,
	})
}

func (worker *WorkerStat) Update(db *gorm.DB) {
	if newDate := time.Now(); worker.LastDate.Day() != newDate.Day() {
		daysBetween := 0
		for {
			newDate.AddDate(0, 0, -1)
			daysBetween++
			if newDate.Format("DDMMYY") == worker.LastDate.Format("DDMMYY") {
				break
			}
		}
		newMessages := worker.Messages[daysBetween:]
		for i := 5 - len(newMessages); i != 0; i-- {
			newMessages = append(newMessages, 0)
		}

		db.Model(&worker).Where("worker = ?", worker.Worker).Updates(&WorkerStat{
			Messages: newMessages,
			LastDate: newDate,
		})

		worker.Messages = newMessages
		worker.LastDate = newDate
	}
}
