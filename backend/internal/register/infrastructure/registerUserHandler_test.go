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

	data := domain.UserRequest{
		Email:    "test@example.com",
		Username: "testuser",
		Name:     "testuser",
		Phone:    "47818189",
		Password: "TestPassword123!",
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

	var response struct {
		Response domain.UserResponse `json:"response"`
		Status   string              `json:"status"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &response)

	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.Response.Email != data.Email {
		t.Errorf("Expected email to be %s, but got %s", data.Email, response.Response.Email)
	}

	if response.Response.Username != data.Username {
		t.Errorf("Expected username to be %s, but got %s", data.Username, response.Response.Username)
	}
}
