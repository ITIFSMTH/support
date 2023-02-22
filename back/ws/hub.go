package ws

import (
	"encoding/json"
	"fmt"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Inbound messages from the clients.
	broadcast chan Request
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

type Request struct {
	client  *Client
	message []byte
}

type Message struct {
	ID     int    `json:"id"`
	Action string `json:"action"`
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

var hub *Hub

func InitHub() {
	hub = &Hub{
		broadcast:  make(chan Request),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func GetHub() *Hub {
	return hub
}

func (h *Hub) RunWS() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case request := <-h.broadcast:
			// Format Message From JSON
			var message Message
			if err := json.Unmarshal(request.message, &message); err != nil {
				request.client.send <- []byte(ErrorBadRequest)
				continue
			}

			fmt.Println(message)

			// Answer on request
			if message.Action == ActionGetChats {
				GetChats(&request)
			} else if message.Action == ActionGetChat {
				GetChat(&request, &message)
			} else if message.Action == ActionSendMessage {
				SendMessage(h, &request, &message)
			} else if message.Action == ActionReadChat {
				ReadChat(h, &request, &message)
			} else {
				request.client.send <- []byte(ErrorBadRequest)
			}
		}
	}
}

func (h *Hub) SendAll(message interface{}) error {
	// Create JSON message
	chatJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for client := range h.clients {
		select {
		case client.send <- chatJson:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}

	return nil
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
