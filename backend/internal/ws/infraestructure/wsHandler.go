package infraestructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/unmsmfisi-socialapplication/social_app/internal/ws/domain"
)

type Handler struct {
	hub *domain.Hub
}

func NewHandler(h *domain.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomReq

	//Reading complete request body
	emptyRequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Transform into request body struct
	err = json.Unmarshal(emptyRequestBody, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.hub.Rooms[req.ID] = &domain.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*domain.Client),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	serializedBody, _ := json.Marshal(&req)
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
