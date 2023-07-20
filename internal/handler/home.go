package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

// updateInfo updates users' data
func (h *Handler) updateInfo(c echo.Context) error {
	user := new(models.User)
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid user id: %v", err)
		c.String(echo.ErrBadRequest.Code, "Invalid user id")
	}

	if err := h.storage.GetUserByID(user, uint(userID)); err != nil {
		log.Printf("No such user in the database: %v", err)
		return c.String(echo.ErrNotFound.Code, "No such user in the database")
	}

	if err := c.Bind(&user); err != nil {
		log.Printf("Incorrect user data: %v", err)
		return c.String(echo.ErrBadRequest.Code, "Incorrect user data")
	}

	if err := h.storage.UpdateUserInfo(user); err != nil {
		log.Printf("Error updating users' info: %v", err)
		c.String(echo.ErrInternalServerError.Code, "Error updating users' info")
	}

	log.Printf("Users' info successfully updated")
	return c.JSON(http.StatusOK, "Your data successfully updated")
}
