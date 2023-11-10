package infrastructure

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
)

type NotificationsHandler struct {
    eventManager *events.EventManager
}

func NewNotificationsHandler(eventManager *events.EventManager) *NotificationsHandler {
	return &NotificationsHandler{eventManager: eventManager}
}

func (nh *NotificationsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Handle error
        conn.Close()
		return
	}

	// Add the WebSocket client to the event manager
	nh.eventManager.AddWSClient(conn)
}
