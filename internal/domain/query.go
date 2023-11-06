package domain

// SingleQuery структура, предназначенная для работы с одним вопросом
type SingleQuery struct {
	Key   string `json:"key"`
	Query string `json:"inquiry"`
}
