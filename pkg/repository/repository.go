package repository

import (
	"gorm.io/gorm"
)

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
