// internal/events/events.go

package events

import "github.com/gorilla/websocket"

// EventListener is an interface that objects wishing to listen to events must implement
type EventListener interface {
	OnEvent(eventName string, eventData interface{})
}

// EventManager manages events and subscribers
type EventManager struct {
	listeners map[string][]EventListener
	WSClients map[*websocket.Conn]bool
}

// NewEventManager creates a new instance of EventManager
func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[string][]EventListener),
		WSClients: make(map[*websocket.Conn]bool),
	}
}

// AddEventListener adds a subscriber to a specific event
func (em *EventManager) AddEventListener(eventName string, listener EventListener) {
	em.listeners[eventName] = append(em.listeners[eventName], listener)
}

// AddWSClient adds a WebSocket client to the list
func (em *EventManager) AddWSClient(client *websocket.Conn) {
	em.WSClients[client] = true

	go func() {
		for {
			defer client.Close()
			_, _, err := client.ReadMessage()
			if err != nil {
				delete(em.WSClients, client)
				break
			}
		}
	}()
}

// TriggerEvent triggers an event to all subscribers
func (em *EventManager) TriggerEvent(eventName string, eventData interface{}) {
	for _, listener := range em.listeners[eventName] {
		listener.OnEvent(eventName, eventData)
	}

	// Send the event to all WebSocket clients
	for client := range em.WSClients {
		err := client.WriteJSON(eventData)
		if err != nil {
			// Manejar el error, por ejemplo, eliminando el cliente WebSocket
			client.Close()
			delete(em.WSClients, client)
		}
	}
}
