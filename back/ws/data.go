package ws

const (
	ActionGetChats    string = "get chats"
	ActionGetChat     string = "get chat"
	ActionReadChat    string = "read chat"
	ActionSendMessage string = "send message"

	CallNewMessage string = "new message"
	CallNewChat    string = "new chat"
	CallReadedChat string = "readed chat"

	ErrorBadRequest         string = "{\"error\": \"bad request\"}"
	ErrorServer             string = "{\"error\": \"server error\"}"
	ErrorChatDontExists     string = "{\"error\": \"chat dont exists\"}"
	ErrorOperatorDontExists string = "{\"error\": \"operator dont exists\"}"
)
