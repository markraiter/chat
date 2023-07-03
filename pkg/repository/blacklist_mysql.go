package repository

import (
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

type BlacklistMySQL struct {
	db *gorm.DB
}

func NewBlacklistMySQL(db *gorm.DB) *BlacklistMySQL {
	return &BlacklistMySQL{db: db}
}

func (r *BlacklistMySQL) AddToBlacklist(blockedUser *models.Blacklist) error {
	r.db.Create(&blockedUser)

	return nil
}

func (r *BlacklistMySQL) RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error {
	r.db.Where("user_id = ? AND blocked_user_id = ?", userID, blockedUserID).Delete(&blockedUser)

	return nil
}
