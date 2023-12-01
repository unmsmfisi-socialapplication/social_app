package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	email := "user@example.com"
	username := "user123"
	password := "password123"

	user, err := NewUser(email, username, password)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if user.Email != email {
		t.Errorf("Expected email to be %s, but got %s", email, user.Email)
	}

	if user.Username != username {
		t.Errorf("Expected username to be %s, but got %s", username, user.Username)
	}

	if user.Password == password {
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
	}

	expectedResponse := UserResponse{
		Username: "testuser",
		Email:    "test@example.com",
	}

	result := UserToUserResponse(user)

	if result != expectedResponse {
		t.Errorf("Expected %v, but got %v", expectedResponse, result)
	}
}

func TestValidateUserRequest(t *testing.T) {
	validUserRequest := UserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	emptyEmailUserRequest := UserRequest{
		Username: "testuser",
		Email:    "",
		Password: "password123",
	}

	emptyUsernameUserRequest := UserRequest{
		Username: "",
		Email:    "test@example.com",
		Password: "password123",
	}

	emptyPasswordUserRequest := UserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "",
	}

	err := ValidateUserRequest(validUserRequest)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	err = ValidateUserRequest(emptyEmailUserRequest)
	if err == nil || err.Error() != "Email is required" {
		t.Errorf("Expected 'Email is required' error, but got %v", err)
	}

	err = ValidateUserRequest(emptyUsernameUserRequest)
	if err == nil || err.Error() != "Username is required" {
		t.Errorf("Expected 'Username is required' error, but got %v", err)
	}

	err = ValidateUserRequest(emptyPasswordUserRequest)
	if err == nil || err.Error() != "Password is required" {
		t.Errorf("Expected 'Password is required' error, but got %v", err)
	}
}
