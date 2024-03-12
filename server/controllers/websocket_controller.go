package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server/models"
	"server/utils"
)

type WebsocketController struct {
	hub *models.Hub
}

func GetWebsocket(h *models.Hub) WebsocketController {
	return WebsocketController{hub: h}
}

func (controller *WebsocketController) CreateRoom(con *gin.Context) {
	var payload models.CreateRoomReq
	err := con.ShouldBindJSON(&payload)

	if err != nil {
		utils.RespondError(con, 400, err)
		return
	}
	rooms := controller.hub.Rooms
	_, exist := rooms[payload.ID]

	if exist {
		utils.RespondError(con, 400, errors.New("room id already exist"))
		return
	}
	rooms[payload.ID] = &models.Room{
		ID:      payload.ID,
		Name:    payload.Name,
		Clients: make(map[string]*models.Client),
	}
	con.JSON(201, rooms[payload.ID])
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (controller *WebsocketController) JoinRoom(con *gin.Context) {
	conn, err := upgrader.Upgrade(con.Writer, con.Request, nil)

	if err != nil {
		utils.RespondError(con, 500, err)
		return
	}
	roomID := con.Param("roomId")
	clientID := con.Query("userId")
	username := con.Query("username")

	client := &models.Client{
		ID:       clientID,
		Message:  make(chan *models.Message, 10),
		Username: username,
		RoomID:   roomID,
		Conn:     conn,
	}

	msg := &models.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	controller.hub.Register <- client
	controller.hub.Broadcast <- msg

	go client.WriteMessage()
	client.ReadMessage(controller.hub)
}

func (controller *WebsocketController) GetRooms(context *gin.Context) {
	rooms := make([]models.RoomRes, 0)

	for _, room := range controller.hub.Rooms {
		rooms = append(rooms, models.RoomRes{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	context.JSON(http.StatusOK, rooms)
}
func (controller *WebsocketController) GetClient(context *gin.Context) {
	var clients []models.ClientRes
	roomId := context.Param("roomId")
	room, exist := controller.hub.Rooms[roomId]

	if !exist {
		clients = make([]models.ClientRes, 0)
		context.JSON(200, clients)
		return
	}
	for _, client := range room.Clients {
		clients = append(clients, models.ClientRes{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	context.JSON(200, clients)

}
