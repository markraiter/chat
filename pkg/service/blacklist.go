package service

import (
	"github.com/markraiter/chat/models"
	"github.com/markraiter/chat/pkg/repository"
)

type Blacklist interface {
	AddToBlacklist(blockedUser *models.Blacklist) error
	RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error
	IsUserBlocked(userID, blockedUserID int) bool
}

type BlacklistService struct {
	repo repository.Blacklist
}

func NewBlacklistService(repo repository.Blacklist) *BlacklistService {
	return &BlacklistService{repo: repo}
}

func (s *BlacklistService) AddToBlacklist(blockedUser *models.Blacklist) error {
	return s.repo.AddToBlacklist(blockedUser)
}

func (s *BlacklistService) RemoveFromBlacklist(userID, blockedUserID int, blockedUser *models.Blacklist) error {
	return s.repo.RemoveFromBlacklist(userID, blockedUserID, blockedUser)
}

func (s *BlacklistService) IsUserBlocked(userID, blockedUserID int) bool {
	blockedUser, err := s.repo.GetBlockedUser(userID, blockedUserID)
	if err != nil {
		return false
	}

	return blockedUser != nil
}
