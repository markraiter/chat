package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

func (h *Handler) addToBlacklist(c echo.Context) error {
	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}
	blockedUserID, err := strconv.Atoi(c.FormValue("blocked_user_id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	blockedUser := models.Blacklist{
		UserID:        userID,
		BlockedUserID: blockedUserID,
	}

	if err := h.services.Blacklist.AddToBlacklist(&blockedUser); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, blockedUser)
}

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
