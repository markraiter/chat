package mysql

import (
	"fmt"

	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

// HomeOperations interface describes all methods for home page
type HomeOperations interface {
	UpdateUserInfo(user *models.User) error
	GetUserByID(user *models.User, id uint) error
}

// Home struct describes Home entity for storage
type Home struct {
	db *gorm.DB
}

// NewHome function is a constructor function for Home struct
func NewHome(db *gorm.DB) *Home {
	return &Home{db: db}
}

// UpdateUserInfo function can update any users' data
func (s *Home) UpdateUserInfo(user *models.User) error {
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	if err := s.db.Save(&user).Error; err != nil {
		return fmt.Errorf("error updating users' info: %w", err)
	}

	return nil
}

// GetUserByID retrieves user by ID
func (s *Auth) GetUserByID(user *models.User, id uint) error {
	if err := s.db.First(&user, id).Error; err != nil {
		return fmt.Errorf("no such user in the database: %w", err)
	}

	return nil
}
