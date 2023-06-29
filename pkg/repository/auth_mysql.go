package repository

import (
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

type AuthMySQL struct {
	db *gorm.DB
}

func NewAuthMySQL(db *gorm.DB) *AuthMySQL {
	return &AuthMySQL{
		db: db,
	}
}

func (r *AuthMySQL) CreateUser(user models.User) (int, error) {
	r.db.Create(&user)

	return int(user.ID), nil
}

func (r *AuthMySQL) GetUser(username, password string) (models.User, error) {
	var user models.User
	r.db.First(&user, username, password)

	return user, nil
}
