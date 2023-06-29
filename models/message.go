package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID   int    `json:"user_id" gorm:"many2many, foreignKey: ID"`
	Username string `json:"username"`
	Body     string `json:"body"`
}
