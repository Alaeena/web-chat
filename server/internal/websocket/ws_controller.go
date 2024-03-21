package websocket

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server/utils"
)

type controller struct {
	hub *Hub
}

func GetController(h *Hub) Controller {
	return &controller{hub: h}
}

func (controller *controller) CreateRoom(con *gin.Context) {
	var payload CreateRoomReq
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
	rooms[payload.ID] = &Room{
		ID:      payload.ID,
		Name:    payload.Name,
		Clients: make(map[string]*Client),
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

func (controller *controller) JoinRoom(con *gin.Context) {
	conn, err := upgrader.Upgrade(con.Writer, con.Request, nil)

	if err != nil {
		utils.RespondError(con, 500, err)
		return
	}
	roomID := con.Param("roomId")
	clientID := con.Query("userId")
	username := con.Query("username")

	client := &Client{
		ID:       clientID,
		Message:  make(chan *Message, 10),
		Username: username,
		RoomID:   roomID,
		Conn:     conn,
	}

	msg := &Message{
		Content:  fmt.Sprintf("%v tham gia vào phòng", client.Username),
		RoomID:   roomID,
		Username: username,
	}

	controller.hub.Register <- client
	controller.hub.Broadcast <- msg

	go client.WriteMessage()
	client.ReadMessage(controller.hub)
}

func (controller *controller) GetRooms(context *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, room := range controller.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	context.JSON(http.StatusOK, rooms)
}
func (controller *controller) GetClient(context *gin.Context) {
	var clients []ClientRes
	roomId := context.Param("roomId")
	room, exist := controller.hub.Rooms[roomId]

	if !exist {
		clients = make([]ClientRes, 0)
		context.JSON(200, clients)
		return
	}
	for _, client := range room.Clients {
		clients = append(clients, ClientRes{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	context.JSON(200, clients)

}
