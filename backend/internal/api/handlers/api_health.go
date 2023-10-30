package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/markraiter/chat/internal/util"
)

var validate = validator.New() //nolint:gochecknoglobals

// @Summary Show the status of server.
// @Description Ping health of API for Docker.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get].
func HandlerAPIHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, util.Response{Message: "healthy"})
	}
}
