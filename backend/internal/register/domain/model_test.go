//go:build unit
// +build unit

package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	userCreate := UserCreate{
		Email:    "user@example.com",
		Username: "user123",
		Password: "Password123!",
	}

	user, err := NewUser(userCreate)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if user.Email != userCreate.Email {
		t.Errorf("Expected email to be %s, but got %s", userCreate.Email, user.Email)
	}

	if user.Username != userCreate.Username {
		t.Errorf("Expected username to be %s, but got %s", userCreate.Username, user.Username)
	}

	if user.Password == userCreate.Password {
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
