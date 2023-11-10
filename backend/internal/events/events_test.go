package events_test

import (
	"testing"

	"github.com/gorilla/websocket"
	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
)

type mockEventListener struct {
	ReceivedEvents map[string]interface{}
}

func NewMockEventListener() *mockEventListener {
	return &mockEventListener{
		ReceivedEvents: make(map[string]interface{}),
	}
}

func (el *mockEventListener) OnEvent(eventName string, eventData interface{}) {
	el.ReceivedEvents[eventName] = eventData
}

func TestEventManager_AddEventListener(t *testing.T) {
	// Setup
	em := events.NewEventManager()
	eventName := "testEvent"
	listener := NewMockEventListener()

	// Add event listener
	em.AddEventListener(eventName, listener)

	// Trigger event
	eventData := "EventData"
	em.TriggerEvent(eventName, eventData)

	// Assert
	if elData, ok := listener.ReceivedEvents[eventName]; !ok || elData != eventData {
		t.Errorf("Unexpected received event data: got %v, expected %v", elData, eventData)
	}
}

func TestEventManager_TriggerEvent_MultipleListeners(t *testing.T) {
	// Setup
	em := events.NewEventManager()
	eventName := "testEvent"
	listener1 := NewMockEventListener()
	listener2 := NewMockEventListener()

	// Add event listeners
	em.AddEventListener(eventName, listener1)
	em.AddEventListener(eventName, listener2)

	// Trigger event
	eventData := "EventData"
	em.TriggerEvent(eventName, eventData)

	// Assert
	if elData, ok := listener1.ReceivedEvents[eventName]; !ok || elData != eventData {
		t.Errorf("Unexpected received event data for listener1: got %v, expected %v", elData, eventData)
	}

	if elData, ok := listener2.ReceivedEvents[eventName]; !ok || elData != eventData {
		t.Errorf("Unexpected received event data for listener2: got %v, expected %v", elData, eventData)
	}
}

func TestEventManager_AddWSClient(t *testing.T) {
	// Setup
	em := events.NewEventManager()
	client := &websocket.Conn{}

	// Add WebSocket client
	em.AddWSClient(client)

	// Trigger event
	eventName := "wsEvent"
	eventData := "WS Data"
	em.TriggerEvent(eventName, eventData)

	// Assert
	if len(em.WSClients) != 1 {
		t.Errorf("Expected one WebSocket client in the manager, got %d", len(em.WSClients))
	}
}
