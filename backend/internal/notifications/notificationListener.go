package notifications

import (
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
)

// NotificationListener is an event listener that handles notifications
type NotificationListener struct {
    eventManager *events.EventManager
}

// NewNotificationListener creates a new instance of NotificationListener
func NewNotificationListener(em *events.EventManager) *NotificationListener {
    return &NotificationListener{
        eventManager: em,
    }
}

// OnEvent implements the EventListener interface to handle events
func (nl *NotificationListener) OnEvent(eventName string, eventData interface{}) {
    // Handle the "postCreated" event and send notifications
    if eventName == "postCreated" {
        nl.SendNotification("New post created!")
    }
}

// SendNotification sends a notification (implement your specific logic here)
func (nl *NotificationListener) SendNotification(message string) {
    fmt.Println("Sending notification: ", message)
    // Implement the logic to send the notification, whether through WebSockets, emails, etc.
}

