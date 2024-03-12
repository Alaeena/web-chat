package router

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/db/database"
	"server/services"
)

func AuthRoute(router *gin.RouterGroup, queries *database.Queries) {
	service := services.GetUser(queries)
	controller := controllers.GetUser(service)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

}
