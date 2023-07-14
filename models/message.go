package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID        int    `json:"user_id" gorm:"many2many, foreignKey: ID"`
	BlockedUserID int    `json:"blocked_user_id"`
	Username      string `json:"username" gorm:"not null"`
	Body          string `json:"body"`
}
