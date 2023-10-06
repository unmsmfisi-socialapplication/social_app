package test

import (
	"testing"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

func TestNewUser(t *testing.T) {
	phone := "123456789"
	email := "user@example.com"
	username := "user123"
	password := "password123"

	user, err := domain.NewUser(phone, email, username, password)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if user.Phone != phone {
		t.Errorf("Expected phone to be %s, but got %s", phone, user.Phone)
	}

	if user.Email != email {
		t.Errorf("Expected email to be %s, but got %s", email, user.Email)
	}

	if user.User_name != username {
		t.Errorf("Expected username to be %s, but got %s", username, user.User_name)
	}

	if user.Password == password {
		t.Errorf("Expected password to be hashed, but got %s", user.Password)
	}

	
}

func TestHashPassword(t *testing.T) {
	password := "password123"

	hashedPassword, err := domain.HashPassword(password)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Errorf("Expected a non-empty hashed password, but got an empty string")
	}

	
}
