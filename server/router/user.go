package router

import (
	"github.com/gin-gonic/gin"
	"server/db/mongo/mongodb"
	"server/internal/user"
	"server/internal/websocket"
)

func AuthRoute(router *gin.RouterGroup, private *gin.RouterGroup, hub *websocket.Hub, queries *mongodb.Queries) {
	service := user.GetMongoService(queries)
	controller := user.GetMongoController(service, hub)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	private.GET("/", controller.Detail)
	private.POST("/logout", controller.Logout)
}
