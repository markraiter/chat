package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Password  string
	Avatar    string
	Friends   []*User
	Blacklist []*User
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []User{}

		if err := db.Find(users).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, users)
	}
}

func (u *User) GetUserByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		user := new(User)

		if err := db.First(user, userID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (u *User) generateToken() (string, error) {
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

func (u *User) Register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, user)
	}
}

func (u *User) Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		var dbUser User

		if err := db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		token, err := u.generateToken()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
