package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markraiter/chat/internal/configs"
	"github.com/markraiter/chat/internal/models"
	"github.com/markraiter/chat/internal/util"
)

type Service interface {
	CreateUser(cfg configs.Config, c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error)
	Login(cfg configs.Config, c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error)
}

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
// @Param input body models.CreateUserReq true "account info"
// @Success 201 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 406 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /signup [post].
func (h *Handler) CreateUser(cfg configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.CreateUserReq
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

			return
		}

		if err := validate.Struct(u); err != nil {
			c.JSON(http.StatusNotAcceptable, util.Response{Message: err.Error()})

			return
		}

		res, err := h.Service.CreateUser(cfg, c.Request.Context(), &u)
		if err != nil {
			if errors.Is(err, util.ErrEmailExist) || errors.Is(err, util.ErrUsernameExist) {
				c.JSON(http.StatusNotAcceptable, util.Response{Message: err.Error()})

				return
			}
			c.JSON(http.StatusInternalServerError, util.Response{Message: err.Error()})

			return
		}

		c.JSON(http.StatusCreated, util.Response{Message: res.ID})
	}
}

// @Summary Login
// @Tags Auth
// @Description Login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.LoginUserReq true "credentials"
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 406 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /login [post].
func (h *Handler) Login(cfg configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.LoginUserReq
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

			return
		}

		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusNotAcceptable, util.Response{Message: err.Error()})

			return
		}

		u, err := h.Service.Login(cfg, c.Request.Context(), &user)
		if err != nil {
			if errors.Is(err, util.ErrWrongCredentials) {
				c.JSON(http.StatusUnauthorized, util.Response{Message: err.Error()})

				return
			}
			c.JSON(http.StatusInternalServerError, util.Response{Message: err.Error()})

			return
		}

		c.SetCookie("jwt", u.AccessToken, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, util.Response{Message: "you are logged in"})
	}
}

// @Summary Logout
// @Tags Auth
// @Description Logout
// @ID logout
// @Produce  json
// @Success 200 {object} util.Response
// @Router /logout [get].
func (h *Handler) Logout(cfg configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("jwt", "", -1, "", "", false, true)
		c.JSON(http.StatusOK, util.Response{Message: "logout successful"})
	}
}
