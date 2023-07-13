package models

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model
	UserID        int `json:"user_id"`
	BlockedUserID int `json:"blocked_user_id"`
}
