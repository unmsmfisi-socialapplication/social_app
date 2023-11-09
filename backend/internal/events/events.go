package events

// EventListener is an interface that objects wishing to listen to events must implement
type EventListener interface {
    OnEvent(eventName string, eventData interface{})
}

// EventManager manages events and subscribers
type EventManager struct {
    listeners map[string][]EventListener
}

// NewEventManager creates a new instance of EventManager
func NewEventManager() *EventManager {
    return &EventManager{
        listeners: make(map[string][]EventListener),
    }
}

// AddEventListener adds a subscriber to a specific event
func (em *EventManager) AddEventListener(eventName string, listener EventListener) {
    em.listeners[eventName] = append(em.listeners[eventName], listener)
}

// TriggerEvent triggers an event to all subscribers
func (em *EventManager) TriggerEvent(eventName string, eventData interface{}) {
    for _, listener := range em.listeners[eventName] {
        listener.OnEvent(eventName, eventData)
    }
}
