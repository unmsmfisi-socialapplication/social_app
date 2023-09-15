package domain

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

// Hub model and constructor
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

// Manages registered, unregistered & broadcast channels
func (h *Hub) RunChatManager() {
	for {
		select {
		case clt := <-h.Register:
			_, exists := h.Rooms[clt.RoomID]
			if exists {
				r := h.Rooms[clt.RoomID]

				_, exists := r.Clients[clt.ID]
				if !exists {
					r.Clients[clt.ID] = clt
				}
			}

		case clt := <-h.Unregister:
			_, exists := h.Rooms[clt.RoomID]
			if exists {
				_, exists := h.Rooms[clt.RoomID].Clients[clt.ID]
				if !exists {
					//broadcast msg for client leaving the room
					if len(h.Rooms[clt.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "A user had left the chat",
							RoomID:   clt.RoomID,
							Username: clt.Username,
						}
					}

					delete(h.Rooms[clt.RoomID].Clients, clt.ID)
					close(clt.Message)
				}
			}

		case msg := <-h.Broadcast:
			_, exists := h.Rooms[msg.RoomID]
			if exists {
				for _, clt := range h.Rooms[msg.RoomID].Clients {
					clt.Message <- msg
				}
			}
		}
	}
}
