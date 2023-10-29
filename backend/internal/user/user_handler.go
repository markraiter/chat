package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markraiter/chat/internal/util"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

// @Summary SignUp
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body user.CreateUserReq true "account info"
// @Success 201 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 406 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /signup [post].
func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Response{Message: err.Error()})

		return
	}

	c.JSON(http.StatusCreated, util.Response{Message: res.ID})

	return
}

// @Summary Login
// @Tags Auth
// @Description Login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body user.LoginUserReq true "account info"
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 406 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /login [post].
func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

		return
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Response{Message: err.Error()})

		return
	}

	c.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, util.Response{Message: "you are logged in"})

	return
}

// @Summary Logout
// @Tags Auth
// @Description Logout
// @ID logout
// @Produce  json
// @Success 200 {object} util.Response
// @Router /logout [get].
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, util.Response{Message: "logout successful"})

	return
}
