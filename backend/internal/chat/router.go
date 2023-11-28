package chat

import (
	"github.com/go-chi/chi"

	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/infraestructure"
)

func ChatModuleRouter(chatHandler *infraestructure.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/createRoom", chatHandler.CreateRoom)
	r.Get("/joinRoom/{roomId}", chatHandler.JoinRoom)
	r.Get("/getRooms", chatHandler.GetRooms)
	r.Get("/getClients/{roomId}", chatHandler.GetClients)

	return r
}
