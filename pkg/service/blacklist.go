package service

import (
	"github.com/markraiter/chat/models"
	"github.com/markraiter/chat/pkg/repository"
)

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
