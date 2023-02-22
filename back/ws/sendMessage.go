package ws

import (
	"fmt"
	"support-back/bot"
	"support-back/db"
	"support-back/models"
	"support-back/shared"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type NewMessage struct {
	Action string `json:"action"`
	ChatID int    `json:"chat_id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func SendMessage(h *Hub, request *Request, message *Message) {
	// Get DB
	db := db.GetDb()

	// Validation
	if len(message.Text) < 1 || len(message.Text) > 4096 {
		request.client.send <- []byte(ErrorBadRequest)
		return
	}

	// Message To Send
	messageDB := &models.Message{
		ChatID: message.ChatID,
		UserID: int64(request.client.operatorId),
		Text:   message.Text,
	}

	// Get Operator
	var operator models.Worker
	db.Model(&models.Worker{}).Where("id = ?", request.client.operatorId).Select("id, login").Scan(&operator)
	if operator.ID == 0 {
		request.client.send <- []byte(ErrorOperatorDontExists)
		return
	}

	// Get Chat
	var chat models.Chat
	db.Model(&models.Chat{}).Where("id = ?", message.ChatID).Scan(&chat)
	if chat.ID == 0 {
		request.client.send <- []byte(ErrorChatDontExists)
		return
	}

	// Create New Message
	db.Model(&models.Message{}).Create(&messageDB)

	// Add message to chat
	db.Model(&models.Chat{}).Where("id = ?", messageDB.ChatID).Update("messages",
		gorm.Expr("messages || ?", pq.Array([]uint{messageDB.ID})),
	)

	// Add operator to chat
	if !contains(chat.Users, messageDB.UserID) {
		db.Model(&models.Chat{}).Where("id = ?", messageDB.ChatID).Update("users",
			gorm.Expr("users || ?", pq.Array([]int64{messageDB.UserID})),
		)
	}

	// Edit operator statistic
	if operator.Role == shared.RoleOperator {
	        models.WorkerStat{}.NewMessage(db, operator.ID)
	}

	// Send message to bot
	bot.SendMessage(chat.UserID, fmt.Sprintf("Новое сообщение от \"%s\":\n%s", operator.Login, message.Text))

	// Send to all
	h.SendAll(&NewMessage{
		Action: CallNewMessage,
		Sender: operator.Login,
		ChatID: int(chat.ID),
		UserID: request.client.operatorId,
		Text:   message.Text,
	})
}
