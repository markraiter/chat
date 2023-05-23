package models

type Message struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Body     string `json:"body"`
}
