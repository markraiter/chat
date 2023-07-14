package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

func (h *Handler) addToFriends(c echo.Context) error {
	friendship := new(models.Friendship)

	if err := c.Bind(&friendship); err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	if err := h.services.FriendList.AddFriend(friendship); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, friendship)
}

// TODO: DEBUG THIS!!!

func (h *Handler) deleteFriend(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.String(echo.ErrBadRequest.Code, "invalid user_id")
	}
	friendID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	friendship := new(models.Friendship)
	if err := c.Bind(&friendship); err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	if err := h.services.FriendList.DeleteFriend(userID, friendID, friendship); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(http.StatusOK, "friend removed successfully")
}
