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
