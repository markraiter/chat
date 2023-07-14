package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

func (h *Handler) addToBlacklist(c echo.Context) error {
	blockedUser := new(models.Blacklist)

	if err := c.Bind(&blockedUser); err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	if err := h.services.Blacklist.AddToBlacklist(blockedUser); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, blockedUser)
}

// TODO: DEBUG THIS!!!

func (h *Handler) deleteFromBlacklist(c echo.Context) error {
	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}
	blockedUserID, err := strconv.Atoi(c.FormValue("blocked_user_id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	if err := h.services.Blacklist.RemoveFromBlacklist(userID, blockedUserID, &models.Blacklist{}); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, "user removed from blacklist successfully")
}
