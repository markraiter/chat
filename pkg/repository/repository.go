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
}

type Blacklist interface {
}

type Repository struct {
	Authorization
	FriendList
	Blacklist
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
	}
}
