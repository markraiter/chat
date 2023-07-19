package mysql

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/markraiter/chat/models"
	"gorm.io/gorm"
)

// Authorization interface describes all methods for authorization
type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(username, password string) (*models.User, error)
	GetUsername(username string) string
	GenerateToken(username, password string) (string, error)
}

// Auth struct describes Auth entity for storage
type Auth struct {
	db *gorm.DB
}

// NewAuth function is a construction function for Auth struct
func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db: db}
}

// CreateUser function creates new user in the database
func (s *Auth) CreateUser(user models.User) (uint, error) {
	if err := user.BeforeCreate(); err != nil {
		return 0, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("error creating user in database: %w", err)
	}

	return user.ID, nil
}

// GetUser function finds user in the database by username
func (s *Auth) GetUser(username, password string) (*models.User, error) {
	user := new(models.User)

	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil || !user.ComparePassword(password) {
		return nil, fmt.Errorf("no such user in the database: %w", err)
	}

	return user, nil
}

// GetUsername retreives username from existing user
func (s *Auth) GetUsername(username string) string {
	user := new(models.User)

	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return ""
	}

	return user.Username
}

// GenerateToken function generates entry token for user to login the chat
func (s *Auth) GenerateToken(username, password string) (string, error) {
	user, err := s.GetUser(username, password)
	if err != nil {
		return "", fmt.Errorf("no such user in the database: %w", err)
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(12 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("error getting tokenstring")
	}

	return tokenString, nil
}
