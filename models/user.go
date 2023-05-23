package models

type User struct {
	ID       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
