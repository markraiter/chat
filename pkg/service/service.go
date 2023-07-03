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
	AddFriend(friendship *models.Friendship) error
	DeleteFriend(userID, friendID int, friendship *models.Friendship) error
}

type Blacklist interface {
	AddToBlacklist(blockedUser *models.Blacklist) error
	RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error
}

type Service struct {
	Authorization
	FriendList
	Blacklist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		FriendList:    NewFriendlistService(repos.FriendList),
		Blacklist:     NewBlacklistService(repos.Blacklist),
	}
}
