package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/websocket"
)

func WebsocketRoute(router *gin.RouterGroup, hub *websocket.Hub) {
	controller := websocket.GetController(hub)

	//HTTP
	router.GET("/", controller.GetRooms)
	router.POST("/", controller.CreateRoom)
	router.GET("/:roomId/clients", controller.GetClient)

	//WEBSOCKET
	router.GET("/:roomId", controller.JoinRoom)
}
