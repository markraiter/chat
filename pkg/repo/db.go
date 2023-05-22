package repo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	Err error
}

func NewDatabase() *Database {
	return &Database{}
}

// Initializing DB
func (db *Database) ConnectToDB() {
	dsn := "root:example@tcp(db:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db.DB, db.Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if db.Err != nil {
		log.Fatalf("error connecting to database: %s/n", db.Err.Error())
	}
}

// Migrations
func (db *Database) MakeMigrations() {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Message{})
}

// Register
func (db *Database) Register(c echo.Context) error {
	user := new(models.User)

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

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

// Login
func (db *Database) Login(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var dbUser models.User

	if err := db.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	mUser := models.NewUser()

	token, err := mUser.GenerateToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

// Get all users
func (db *Database) GetUsers(c echo.Context) error {
	users := []models.User{}

	if err := db.DB.Find(users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// Get user by ID
func (db *Database) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user := new(models.User)

	if err := db.DB.First(user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
