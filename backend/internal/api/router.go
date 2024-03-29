package api

import (
	"github.com/gin-gonic/gin"
	"github.com/markraiter/chat/internal/api/handlers"
	"github.com/markraiter/chat/internal/configs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func InitRoutes(cfg configs.Config, userHandler *handlers.Handler, wsHandler *handlers.WSHandler) {
	r = gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", handlers.HandlerAPIHealth())

	r.POST("/signup", userHandler.CreateUser(cfg))
	r.POST("/login", userHandler.Login(cfg))
	r.GET("/logout", userHandler.Logout(cfg))

	ws := r.Group("/ws")
	{
		ws.POST("/create-room", wsHandler.CreateRoom)
		ws.GET("/join-room/:room_id", wsHandler.JoinRoom)
		ws.GET("/get-rooms", wsHandler.GetRooms)
		ws.GET("/get-clients/:room_id", wsHandler.GetClients)
	}

}

func Start(cfg configs.Server) error {
	return r.Run(cfg.AppAddress)
}
