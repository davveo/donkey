package fake

type MessageUnread struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    MessageUnreadData `json:"data"`
}
type MessageUnreadData struct {
	Num0  int `json:"0"`
	Num1  int `json:"1"`
	Num2  int `json:"2"`
	Num3  int `json:"3"`
	Total int `json:"total"`
}
