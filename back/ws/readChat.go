package ws

import (
	"support-back/db"
	"support-back/models"
)

type ReadedChat struct {
	Action string `json:"action"`
	ChatID int    `json:"chat_id"`
}

func ReadChat(h *Hub, request *Request, message *Message) {
	// Get DB
	db := db.GetDb()

	// Validation
	if message.ChatID == 0 {
		request.client.send <- []byte(ErrorBadRequest)
		return
	}

	// Get Chat
	var chat models.Chat
	db.Model(&models.Chat{}).Where("id = ?", message.ChatID).Scan(&chat)
	if chat.ID == 0 {
		request.client.send <- []byte(ErrorChatDontExists)
		return
	}

	// Set Messages Readed = true
	for _, messageId := range chat.Messages {
		db.Model(&models.Message{}).Where("id = ?", int(messageId)).Update("readed", true)
	}

	// Send to all
	h.SendAll(&ReadedChat{
		Action: CallReadedChat,
		ChatID: int(chat.ID),
	})
}
