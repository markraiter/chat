package service

import (
	"github.com/markraiter/chat/pkg/repository"
)

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
