package repository

import (
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type FriendList interface {
	AddFriend(friendship *models.Friendship) error
	DeleteFriend(userID, friendID int, friendship *models.Friendship) error
}

type Blacklist interface {
	AddToBlacklist(blockedUser *models.Blacklist) error
	RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error
}

type Repository struct {
	Authorization
	FriendList
	Blacklist
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
		FriendList:    NewFriendlistMySQL(db),
		Blacklist:     NewBlacklistMySQL(db),
	}
}
