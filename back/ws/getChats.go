package ws

import (
	"encoding/json"
	"support-back/db"
	"support-back/models"
	"time"
)

type ChatsChat struct {
	User          models.User     `json:"user"`
	Operators     []models.Worker `json:"operators"`
	UnreadedCount int             `json:"unreaded_count"`
	ChatID        int             `json:"chat_id"`
	LastMessage   time.Time       `json:"last_message"`
}

type ChatsAnswer struct {
	Action string      `json:"action"`
	Chats  []ChatsChat `json:"chats"`
}

func GetChats(request *Request) {
	// Answer Struct
	answer := &ChatsAnswer{
		Action: ActionGetChats,
	}

	// Get DB
	db := db.GetDb()

	// Get chats from DB
	var chats []models.Chat
	db.Model(&models.Chat{}).Where("messages != '{}'").Find(&chats)

	// Parse chats
	for _, chat := range chats {
		var user models.User
		var operators []models.Worker = []models.Worker{}
		var unreaded int
		var lastMessage time.Time

		// Get User
		db.Model(&models.User{}).Where("tg_id = ?", chat.UserID).Scan(&user)

		// Get Operators
		for _, operatorId := range chat.Users {
			if operatorId < 10000 {
				var operator models.Worker
				db.Model(&models.Worker{}).Where("id = ?", operatorId).Select("Login").Scan(&operator)
				operators = append(operators, operator)
			}
		}

		// Get unreaded count && last message date
		for messageIdx, messageId := range chat.Messages {
			var message models.Message
			db.Model(&models.Message{}).Where("id = ?", messageId).Select("readed, created_at").Scan(&message)
			if !message.Readed {
				unreaded++
			}
			if messageIdx == len(chat.Messages) - 1 {
				lastMessage = message.CreatedAt
			}
		}

		answer.Chats = append(answer.Chats, ChatsChat{
			User:          user,
			Operators:     operators,
			UnreadedCount: unreaded,
			ChatID:        int(chat.ID),
			LastMessage:   lastMessage,
		})
	}

	chatsJson, err := json.Marshal(answer)
	if err != nil {
		request.client.send <- []byte(ErrorServer)
		return
	}

	request.client.send <- []byte(chatsJson)
}
