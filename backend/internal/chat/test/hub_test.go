package test

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/domain"
)

func TestNewHub(t *testing.T) {
	hub := domain.NewHub()

	if hub == nil {
		t.Error("Expected hub to be created, but it was not")
	}

	if len(hub.Rooms) != 0 {
		t.Errorf("Expected hub to have 0 rooms, got %d", len(hub.Rooms))
	}

	select {
	case <-hub.Register:
		t.Error("Expected hub.Register channel to be unbuffered")
	default:
	}

	select {
	case <-hub.Unregister:
		t.Error("Expected hub.Unregister channel to be unbuffered")
	default:
	}

	select {
	case <-hub.Broadcast:
		t.Error("Expected hub.Broadcast channel to be unbuffered")
	default:
	}
}
