package bot

import (
	"errors"
	"support-back/db"
	"support-back/models"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CheckToken(token string) (string, error) {
	testBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return "", err
	}

	return testBot.Self.UserName, nil
}

func SendMessage(userId int64, text string) error {
	// Get DB
	db := db.GetDb()

	// Get User
	var user models.User
	db.Model(&user).Where("tg_id = ?", userId).Scan(&user)

	// Get TG bot
	var bot models.Bot
	db.Model(&bot).Where("id = ?", user.BotID).Scan(&bot)

	if bot.ID == 0 {
		return errors.New("Bot Deleted")
	}

	// Get bot
	tgBot, err := tgbotapi.NewBotAPI(bot.TGAPI)
	if err != nil {
		return err
	}

	// Send Message
	tgBot.Send(tgbotapi.NewMessage(userId, text))
	return nil
}

func SendMailing(text string) {
	type UserWithBot struct {
		TGID  int64
		TGAPI string
	}

	// Get DB
	db := db.GetDb()

	// Get users with bots
	var usersWithBots []UserWithBot
	db.Table("users").Joins("JOIN bots ON users.bot_id = bots.id").Select("users.tg_id, bots.tg_api").Scan(&usersWithBots)

	// Send message to every user
	for _, userWithBot := range usersWithBots {
		if userWithBot.TGAPI == "" {
			continue
		}

		// Get bot
		tgBot, err := tgbotapi.NewBotAPI(userWithBot.TGAPI)
		if err != nil {
			return
		}

		// Send Message
		tgBot.Send(tgbotapi.NewMessage(userWithBot.TGID, text))

		// Avoid telegram rate limits
		time.Sleep(time.Second / 20)
	}
}
