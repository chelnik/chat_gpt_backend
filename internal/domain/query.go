package domain

// SingleQuery структура, предназначенная для работы с одним вопросом
type SingleQuery struct {
	Key     string `json:"key"`
	Inquiry string `json:"inquiry"`
}
