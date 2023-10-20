package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func (m *mockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (m *mockUserRepository) InsertUser(newUser *domain.User) (*domain.User, error) {
	if _, exists := m.users[newUser.Email]; exists {
		return nil, errors.New("User already exists")
	}
	m.users[newUser.Email] = newUser
	return newUser, nil
}

func TestRegisterUserHandler_RegisterUser(t *testing.T) {

	mockUserRepository := &mockUserRepository{
		users: make(map[string]*domain.User),
	}

	mockUseCase := application.NewRegistrationUseCase(mockUserRepository)

	handler := NewRegisterUserHandler(mockUseCase)

	data := map[string]string{
		"phone":     "123456789",
		"email":     "test@example.com",
		"user_name": "testuser",
		"password":  "TestPassword123!",
	}
	requestData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", "/register", bytes.NewReader(requestData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.RegisterUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["email"] != data["email"] {
		t.Errorf("Expected email to be %s, but got %s", data["email"], response["email"])
	}

	if response["user_name"] != data["user_name"] {
		t.Errorf("Expected user_name to be %s, but got %s", data["user_name"], response["user_name"])
	}
}
