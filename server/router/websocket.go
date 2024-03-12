package router

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/db/database"
	"server/models"
)

func WebsocketRoute(router *gin.RouterGroup, queries *database.Queries) {
	hub := models.NewHub()
	controller := controllers.GetWebsocket(&hub)
	go hub.Run()

	//HTTP
	router.GET("/", controller.GetRooms)
	router.POST("/", controller.CreateRoom)
	router.GET("/:roomId/clients", controller.GetClient)

	//WEBSOCKET
	router.GET("/:roomId", controller.JoinRoom)
}
