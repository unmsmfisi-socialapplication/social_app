package infraestructure

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/domain"
)

type Handler struct {
	hub *domain.Hub
}

func NewHandler(h *domain.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var res roomRes

	//Reading complete request body
	emptyRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Transform into request body struct
	err = json.Unmarshal(emptyRequestBody, &res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.hub.Rooms[res.ID] = &domain.Room{
		ID:      res.ID,
		Name:    res.Name,
		Clients: make(map[string]*domain.Client),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	serializedBody, _ := json.Marshal(&res)
	_, _ = w.Write(serializedBody)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//	JoinUrlParams
//
// -> ws/JoinRoom/{roomId}?userId=[userid]&username=[username]
func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	roomID := chi.URLParam(r, "roomId")
	clientID := r.URL.Query().Get("userId")
	username := r.URL.Query().Get("username")

	clt := &domain.Client{
		Conn:     conn,
		Message:  make(chan *domain.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	//initial Message
	msg := &domain.Message{
		Content:  "New user has been connected",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- clt
	h.hub.Broadcast <- msg

	go clt.WriteMessage()
	clt.ReadMessage(h.hub)

}

type roomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]roomRes, 0)

	for _, ro := range h.hub.Rooms {
		rooms = append(rooms, roomRes{
			ID:   ro.ID,
			Name: ro.Name,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	serializedBody, _ := json.Marshal(&rooms)
	_, _ = w.Write(serializedBody)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []ClientRes

	roomID := chi.URLParam(r, "roomId")

	_, exists := h.hub.Rooms[roomID]
	if !exists {
		clients = make([]ClientRes, 0)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		serializedBody, _ := json.Marshal(&clients)
		_, _ = w.Write(serializedBody)
	}

	for _, clt := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientRes{
			ID:       clt.ID,
			Username: clt.Username,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	serializedBody, _ := json.Marshal(&clients)
	_, _ = w.Write(serializedBody)

}
