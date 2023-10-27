//go:build unit
// +build unit

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
