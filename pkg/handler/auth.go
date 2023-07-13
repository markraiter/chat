package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) register(c echo.Context) error {
	var input models.User

	if err := c.Bind(&input); err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) login(c echo.Context) error {
	var input LoginInput

	if err := c.Bind(&input); err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
