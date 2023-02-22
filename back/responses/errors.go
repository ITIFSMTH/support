package responses

const (
	// Global errors
	ErrorBadData = "Невалидные данные, проверьте поля ввода"
	ErrorServer  = "Серверная ошибка, сообщите администратору"
	ErrorAuth    = "Нет прав на выполнение запроса"

	// Admin Errors
	ErrorBotExist           = "Такой бот уже существует"
	ErrorOperatorLoginExist = "Оператор с таким логином уже существует"
	ErrorTokenInvalid       = "Невалидный токен"

	// Auth Errors
	ErrorWrongData = "Неверные данные"
)
