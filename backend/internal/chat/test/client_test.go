package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/domain"
)

func TestClientReadWrite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Error upgrading connection: %v", err)
		}
		defer conn.Close()

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if messageType == websocket.TextMessage {
				if err := conn.WriteMessage(messageType, p); err != nil {
					t.Fatalf("Error writing message: %v", err)
				}
			}
		}
	}))
	defer server.Close()

	clientConn, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatalf("Error connecting to WebSocket server: %v", err)
	}
	defer clientConn.Close()

	client := &domain.Client{
		Conn:     clientConn,
		Message:  make(chan *domain.Message, 10),
		ID:       "testClient",
		RoomID:   "testRoom",
		Username: "testUser",
	}

	hub := domain.NewHub()

	go client.ReadMessage(hub)

	messageToSend := &domain.Message{
		Content:  "Hello, server!",
		RoomID:   "testRoom",
		Username: "testUser",
	}
	client.Message <- messageToSend

	select {
	case receivedMessage := <-client.Message:
		if receivedMessage.Content != messageToSend.Content {
			t.Errorf("Expected received message content to be '%s', got '%s'", messageToSend.Content, receivedMessage.Content)
		}
	}
}
