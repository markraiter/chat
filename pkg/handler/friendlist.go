package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

func (h *Handler) addToFriends(c echo.Context) error {
	userID := c.Get("user_id").(int)
	friendID, err := strconv.Atoi(c.FormValue("friend_id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	friendship := models.Friendship{
		UserID:   userID,
		FriendID: friendID,
	}

	if err := h.services.FriendList.AddFriend(&friendship); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, friendship)
}

func (h *Handler) deleteFriend(c echo.Context) error {
	userID := c.Get("user_id").(int)
	friendID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	if err := h.services.FriendList.DeleteFriend(userID, friendID, &models.Friendship{}); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, "friend removed successfully")
}
