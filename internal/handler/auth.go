package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

// LoginInput is a data type for user's login input
type LoginInput struct {
	Username string
	Password string
}

// signUp gets user's data from request and creates new user in database
func (h *Handler) signUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		log.Printf("Incorrect user data: %v", err)
		return c.String(echo.ErrBadRequest.Code, "Incorrect user data")
	}

	username := h.storage.GetUsername(user.Username)
	if user.Username == username {
		log.Printf("This user is already exists: %v", username)
		return c.String(http.StatusAccepted, "This user is already exists")
	}

	id, err := h.storage.CreateUser(*user)
	if err != nil {
		log.Printf("Error creating user in the database: %v", err)
		return c.String(echo.ErrInternalServerError.Code, "Error creting user in database")
	}

	log.Printf("Successfully created new user by id: %v", id)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// signIn gets login input from user, finds this user in database and returns entry token
func (h *Handler) signIn(c echo.Context) error {
	var input LoginInput

	if err := c.Bind(&input); err != nil {
		log.Printf("Incorrect input: %v", input)
		return c.String(echo.ErrBadRequest.Code, "Incorrect input")
	}

	token, err := h.storage.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Printf("Incorrect username or password: %v", err)
		return c.String(echo.ErrNotFound.Code, "Incorrect username or password")
	}

	log.Printf("Successfully got token: %s", token)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
