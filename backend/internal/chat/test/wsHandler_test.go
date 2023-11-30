package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/chat/infraestructure"
)

func TestCreateRoom(t *testing.T) {
	hub := domain.NewHub()
	handler := infraestructure.NewHandler(hub)

	requestBody := roomRes{
		ID:   "room1",
		Name: "Testing-Room",
	}

	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/create-room", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateRoom(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	room := hub.Rooms["room1"]
	if room == nil {
		t.Error("Expected room to be created, but it was not")
	} else if room.Name != "Testing-Room" {
		t.Errorf("Expected room name to be 'Testing-Room', got %s", room.Name)
	}
}

type roomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestGetRooms(t *testing.T) {
	hub := domain.NewHub()
	hub.Rooms["room1"] = &domain.Room{
		ID:      "room1",
		Name:    "Testing-Room",
		Clients: make(map[string]*domain.Client),
	}
	handler := infraestructure.NewHandler(hub)

	req := httptest.NewRequest("GET", "/get-rooms", nil)
	rec := httptest.NewRecorder()

	handler.GetRooms(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var response []roomRes
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response JSON: %v", err)
	}

	expectedResponse := []roomRes{
		{
			ID:   "room1",
			Name: "Testing-Room",
		},
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}
}
