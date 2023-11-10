package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/events"
	"github.com/unmsmfisi-socialapplication/social_app/internal/notifications/domain"
)

type NotificationListener struct {
    eventManager *events.EventManager
}

func NewNotificationListener(em *events.EventManager) *NotificationListener {
    return &NotificationListener{
        eventManager: em,
    }
}

func (nl *NotificationListener) OnEvent(eventName string, eventData interface{}) {
    if eventName == "postCreated" {
        nl.SendNotification("New post created!", eventData)
    }
}

func (nl *NotificationListener) SendNotification(message string, data interface{}) {
    res := domain.Notification{
        Message: message,
        Data:    data,
    }

    for client := range nl.eventManager.WSClients {
        err := client.WriteJSON(res)
        if err != nil {
            client.Close()
        }
    }
}

