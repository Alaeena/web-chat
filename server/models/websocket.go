package models

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewHub() Hub {
	return Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			room, exist := hub.Rooms[client.RoomID]
			if exist {
				_, exist := room.Clients[client.ID]
				if !exist {
					room.Clients[client.ID] = client
				}
			}
		case client := <-hub.Unregister:
			room, exist := hub.Rooms[client.RoomID]
			if exist {
				_, exist := room.Clients[client.ID]
				if exist {
					delete(room.Clients, client.ID)
					close(client.Message)
				}
			}
		case msg := <-hub.Broadcast:
			room, exist := hub.Rooms[msg.RoomID]
			if exist {
				for _, client := range room.Clients {
					client.Message <- msg
				}
			}
		}
	}
}
