package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"server/db/mongo"
	"server/db/scylla"
	"server/internal/websocket"
	"server/middleware"
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
	port := os.Getenv("PORT")

	scyllaQueries := scylla.Queries()
	mongoQueries := mongo.Queries()
	hub := websocket.NewHub()

	go hub.Run(scyllaQueries)
	handleCORS(router)

	private := router.Group("/auth", middleware.AuthHandler())
	AuthRoute(router.Group("/auth"), private, hub, mongoQueries)
	WebsocketRoute(router.Group("/websocket"), hub)

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "index page"})
	})
	return router.Run("localhost:" + port)
}
