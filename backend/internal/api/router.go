package api

import (
	"github.com/gin-gonic/gin"
	"github.com/markraiter/chat/internal/api/handlers"
	"github.com/markraiter/chat/internal/configs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func InitRouter(cfg configs.Config, userHandler *handlers.Handler, wsHandler *handlers.WSHandler) {
	r = gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", handlers.HandlerAPIHealth())

	r.POST("/signup", userHandler.CreateUser(cfg))
	r.POST("/login", userHandler.Login(cfg))
	r.GET("/logout", userHandler.Logout(cfg))

	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:room_id", wsHandler.JoinRoom)
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients/:room_id", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
