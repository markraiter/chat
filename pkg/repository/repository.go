package repository

import (
	"database/sql"

	"github.com/markraiter/chat/models"
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

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
	}
}
