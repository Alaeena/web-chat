package websocket

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateRoom(con *gin.Context)
	JoinRoom(con *gin.Context)
	GetRooms(context *gin.Context)
	GetClient(context *gin.Context)
}
