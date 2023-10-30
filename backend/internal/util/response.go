package util

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message" example:"response message"`
}

type Responser struct {
	Log *slog.Logger
}

func (r *Responser) Error(c *gin.Context, msg string, status int, err error) {
	errString := fmt.Sprintf("%s error: %s", msg, err.Error())

	r.Log.Error("error_log:", "message", errString)

	c.JSON(status, gin.H{"error": errString})
}

func (r *Responser) Warn(c *gin.Context, msg string, status int, err error) {
	errString := fmt.Sprintf("%s error: %s", msg, err.Error())

	r.Log.Warn("error_log:", "message", errString)

	c.JSON(status, gin.H{"error": errString})
}

func (r *Responser) Info(c *gin.Context, msg string, status int, err error) {
	errString := fmt.Sprintf("%s error: %s", msg, err.Error())

	r.Log.Info("error_log:", "message", errString)

	c.JSON(status, gin.H{"message": errString})
}

func (r *Responser) Debug(c *gin.Context, msg string, status int, err error) {
	errString := fmt.Sprintf("%s error: %s", msg, err.Error())

	r.Log.Debug("error_log:", "message", errString)

	c.JSON(status, gin.H{"message": errString})
}
