package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"server/db"
	"time"
)

func handleCORS(router *gin.Engine) {
	clientUrl := os.Getenv("CLIENT_URL")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{clientUrl},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == clientUrl
		},
		MaxAge: time.Hour,
	}))
}
func Listen() error {
	router := gin.Default()
	queries := db.Client()
	port := os.Getenv("PORT")

	handleCORS(router)
	AuthRoute(router.Group("/auth"), queries)
	WebsocketRoute(router.Group("/websocket"), queries)
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "index page"})
	})

	return router.Run("localhost:" + port)
}
