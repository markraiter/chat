package repository

import (
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

type FriendList interface {
	AddFriend(friendship *models.Friendship) error
	DeleteFriend(userID, friendID int, friendship *models.Friendship) error
}

type FriendlistMySQL struct {
	db *gorm.DB
}

func NewFriendlistMySQL(db *gorm.DB) *FriendlistMySQL {
	return &FriendlistMySQL{db: db}
}

func (r *FriendlistMySQL) AddFriend(friendship *models.Friendship) error {
	r.db.Create(&friendship)

	return nil
}

func (r *FriendlistMySQL) DeleteFriend(userID, friendID int, friendship *models.Friendship) error {
	r.db.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&friendship)

	return nil
}
