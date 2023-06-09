package models

type Message struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Body     string `json:"body"`
}
