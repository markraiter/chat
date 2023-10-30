package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/markraiter/chat/internal/configs"
	"github.com/markraiter/chat/internal/models"
	"github.com/markraiter/chat/internal/util"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) *service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(cfg configs.Config, c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("user_service CreateUser() error: %w", err)
	}

	u := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("user_service CreateUser() error: %w", err)
	}

	res := &models.CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(cfg configs.Config, c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("user_service Login() error: %w", err)
	}

	if err := util.CheckPassword(req.Password, u.Password); err != nil {
		return nil, fmt.Errorf("user_service Login() error: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.Auth.RefreshTokenTTL)),
		},
	})

	ss, err := token.SignedString([]byte(cfg.Auth.SigningKey))
	if err != nil {
		return nil, fmt.Errorf("user_service Login() error: %w", err)
	}

	return &models.LoginUserRes{
		AccessToken: ss,
		ID:          strconv.Itoa(int(u.ID)),
		Username:    u.Username,
	}, nil
}
