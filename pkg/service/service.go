package service

import (
	"github.com/markraiter/chat/models"
	"github.com/markraiter/chat/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type FriendList interface {
}

type Blacklist interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
