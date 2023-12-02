package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {

	userReq := UserRequest{
		Email:    "user@example.com",
		Username: "user123",
		Password: "password123",
	}

	user, err := NewUser(userReq)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if user.Email != userReq.Email {
		t.Errorf("Expected email to be %s, but got %s", userReq.Email, user.Email)
	}

	if user.Username != userReq.Username {
		t.Errorf("Expected username to be %s, but got %s", userReq.Username, user.Username)
	}

	if user.Password == userReq.Password {
		t.Errorf("Expected password to be hashed, but got %s", user.Password)
	}

}

func TestHashPassword(t *testing.T) {
	password := "password123"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Errorf("Expected a non-empty hashed password, but got an empty string")
	}

}

func TestUserToUserResponse(t *testing.T) {
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Phone:    "48941651",
		Photo:    "https://promotonews.com/wp-content/uploads/2020/09/IMG-LOGO.jpg",
		Name:     "asdasd",
	}

	expectedResponse := UserResponse{
		Username: "testuser",
		Email:    "test@example.com",
		Phone:    "48941651",
		Photo:    "https://promotonews.com/wp-content/uploads/2020/09/IMG-LOGO.jpg",
		Name:     "asdasd",
		Type:     "user",
		Role:     1,
		Summary:  "I am a developer",
	}

	result := UserToUserResponse(user)

	if result.Username != expectedResponse.Username ||
		result.Email != expectedResponse.Email ||
		result.Phone != expectedResponse.Phone ||
		result.Photo != expectedResponse.Photo ||
		result.Name != expectedResponse.Name ||
		result.Type != expectedResponse.Type ||
		result.Role != expectedResponse.Role ||
		result.Summary != expectedResponse.Summary {
		t.Errorf("Expected %+v, but got %+v", expectedResponse, result)
	}
}
