package models

import "gorm.io/gorm"

type Friendship struct {
	gorm.Model
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
