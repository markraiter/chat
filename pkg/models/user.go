package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Password  string
	Avatar    string
	Friends   []*User `gorm:"many2many:friendships"`
	Blacklist []*User `gorm:"many2many:blaklist"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"password": u.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *User) AddToBlacklist(blocked *User) {
	u.Blacklist = append(u.Blacklist, blocked)
}

func (u *User) AddToFriends(friend *User) {
	u.Friends = append(u.Friends, friend)
}
