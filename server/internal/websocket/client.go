package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Message struct {
	Content   string    `json:"content"`
	RoomID    string    `json:"roomId"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"Id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (client *Client) WriteMessage() {
	defer func() {
		client.Conn.Close()
	}()

	for {
		msg, ok := <-client.Message
		if !ok {
			return
		}
		client.Conn.WriteJSON(msg)
	}
}
func (client *Client) ReadMessage(hub *Hub) {
	defer func() {
		client.Conn.Close()
		hub.Unregister <- client
		hub.Broadcast <- &Message{
			Content:  fmt.Sprintf("%v đã rời khỏi phòng", client.Username),
			RoomID:   client.RoomID,
			Username: client.Username,
		}
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			isCloseErr := websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			if isCloseErr {
				log.Printf("error: %v\n", err)
			}
			break
		}
		message := &Message{
			Content:  string(msg),
			RoomID:   client.RoomID,
			Username: client.Username,
		}
		hub.Broadcast <- message
	}
}
