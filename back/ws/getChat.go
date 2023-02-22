package ws

import (
	"encoding/json"
	"support-back/db"
	"support-back/models"
)

type ChatMessage struct {
	Message    models.Message `json:"message"`
	SenderName string         `json:"sender_name"`
}

type ChatAnswer struct {
	Action   string        `json:"action"`
	User     models.User   `json:"user"`
	Messages []ChatMessage `json:"messages"`
	ChatID   int           `json:"chat_id"`
}

func GetChat(request *Request, message *Message) {
	// Get DB
	db := db.GetDb()

	// Declare variable
	var chat models.Chat
	var user models.User
	var messages []ChatMessage

	// Get Chat
	db.Model(&models.Chat{}).Where("id = ?", message.ChatID).Scan(&chat)

	// If chat don't exists return
	if chat.UserID == 0 {
		request.client.send <- []byte(ErrorChatDontExists)
		return
	}

	// Get User
	db.Model(&models.User{}).Where("tg_id = ?", chat.UserID).Scan(&user)

	// Get Messages
	for _, messageId := range chat.Messages {
		var message models.Message
		var senderName string
		db.Model(&models.Message{}).Where("id = ?", int(messageId)).Scan(&message)

		if message.UserID > 10000 {
			var user models.User
			db.Model(&models.User{}).Where("tg_id = ?", message.UserID).Select("username").Scan(&user)
			senderName = user.Username
		} else {
			var operator models.Worker
			db.Model(&models.Worker{}).Where("id = ?", message.UserID).Select("login").Scan(&operator)
			senderName = operator.Login
		}

		messages = append(messages, ChatMessage{
			Message:    message,
			SenderName: senderName,
		})
	}

	// Create JSON answer
	chatJson, err := json.Marshal(&ChatAnswer{
		Action:   ActionGetChat,
		User:     user,
		Messages: messages,
		ChatID:   int(chat.ID),
	})
	if err != nil {
		request.client.send <- []byte(ErrorServer)
		return
	}

	request.client.send <- []byte(chatJson)
}
