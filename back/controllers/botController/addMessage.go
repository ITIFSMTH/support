package botController

import (
	"fmt"
	"net/http"
	"support-back/db"
	"support-back/models"
	"support-back/responses"
	"support-back/ws"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type AddMessageRequestBody struct {
	UserID   int64  `json:"user_id" binding:"required"`
	Username string `json:"user_name"`
	Text     string `json:"text" binding:"required"`
	TGAPI    string `json:"tg_api" binding:"required"`
}

type AddChatWSBroadcast struct {
	Action string      `json:"action"`
	ChatID int         `json:"chat_id"`
	User   models.User `json:"user"`
}

func (*BotController) AddMessage(c *gin.Context) {
	// Get Data
	var requestBody AddMessageRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDb()

	// Get Chat
	var chat models.Chat
	db.Model(&models.Chat{}).Where("user_id = ?", requestBody.UserID).Select("messages, id").Scan(&chat)

	// Message
	message := &models.Message{
		ChatID: int64(chat.ID),
		UserID: requestBody.UserID,
		Text:   requestBody.Text,
		Readed: false,
	}

	// Create New Message
	db.Model(&message).Create(&message).Update("readed", false)

	// Add message to chat
	db.Model(&models.Chat{}).Where("id = ?", message.ChatID).Update("messages",
		gorm.Expr("messages || ?", pq.Array([]uint{message.ID})),
	)

	// Get User
	var user models.User
	db.Model(&user).Where("tg_id = ?", requestBody.UserID).Select("username, bot_id, tg_link").Scan(&user)

	// Check is new username?
	if user.Username != requestBody.Username {
		db.Model(&user).Where("tg_id = ?", requestBody.UserID).Update(&models.User{
			TGLink: requestBody.Username,
		})
	}

	// Check is new bot?
	var bot models.Bot
	db.Model(&bot).Where("id = ?", user.BotID).Scan(&bot)
	fmt.Print(bot.TGAPI, requestBody.TGAPI)
	if requestBody.TGAPI != bot.TGAPI {
		var newBot models.Bot
		db.Model(&newBot).Where("tg_api = ?", requestBody.TGAPI).Scan(&newBot)
		db.Model(&user).Update(&models.User{
			BotID: int(newBot.ID),
		})
	}

	h := ws.GetHub()

	// Chat Is New?
	if len(chat.Messages) == 0 {
		h.SendAll(&AddChatWSBroadcast{
			Action: ws.CallNewChat,
			ChatID: int(chat.ID),
			User:   user,
		})
	}

	h.SendAll(&ws.NewMessage{
		Action: ws.CallNewMessage,
		ChatID: int(chat.ID),
		Sender: user.Username,
		Text:   message.Text,
		UserID: int(requestBody.UserID),
	})

	// Send Answer
	c.JSON(http.StatusOK, &responses.Response{})
}
