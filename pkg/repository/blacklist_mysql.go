package repository

import (
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

type Blacklist interface {
	AddToBlacklist(blockedUser *models.Blacklist) error
	RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error
	GetBlockedUser(userID, blockedUserID int) (*models.Blacklist, error)
}

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

func (r *BlacklistMySQL) GetBlockedUser(userID, blockedUserID int) (*models.Blacklist, error) {
	var blockedUser models.Blacklist
	err := r.db.Where("user_id = ? AND blocked_user_id = ?", userID, blockedUserID).First(&blockedUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &blockedUser, nil
}
