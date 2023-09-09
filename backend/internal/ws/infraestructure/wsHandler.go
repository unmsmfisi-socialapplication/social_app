package infraestructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	/*
		h.hub.Rooms[req.ID] = &Room{
			ID:      req.ID,
			Name:    req.Name,
			Clients: make(map[string]*Client),
		}
	*/

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	serializedBody, _ := json.Marshal(&req)
	_, _ = w.Write(serializedBody)
}
