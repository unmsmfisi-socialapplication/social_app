package events

import "github.com/gorilla/websocket"

type EventListener interface {
	OnEvent(eventName string, eventData interface{})
}

type EventManager struct {
	listeners map[string][]EventListener
	WSClients map[*websocket.Conn]bool
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[string][]EventListener),
		WSClients: make(map[*websocket.Conn]bool),
	}
}

func (em *EventManager) AddEventListener(eventName string, listener EventListener) {
	em.listeners[eventName] = append(em.listeners[eventName], listener)
}

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

func (em *EventManager) TriggerEvent(eventName string, eventData interface{}) {
	for _, listener := range em.listeners[eventName] {
		listener.OnEvent(eventName, eventData)
	}
}
