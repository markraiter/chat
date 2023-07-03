package service

import (
	"github.com/markraiter/chat/models"
	"github.com/markraiter/chat/pkg/repository"
)

type FriendlistService struct {
	repo repository.FriendList
}

func NewFriendlistService(repo repository.FriendList) *FriendlistService {
	return &FriendlistService{repo: repo}
}

func (s *FriendlistService) AddFriend(friendship *models.Friendship) error {
	return s.repo.AddFriend(friendship)
}

func (s *FriendlistService) DeleteFriend(userID, friendID int, friendship *models.Friendship) error {
	return s.repo.DeleteFriend(userID, friendID, friendship)
}
