package application

import (
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

func TestAuthenticate(t *testing.T) {
	email := "user@example.com"
	username := "user123"
	password := "Password123!"

    user, _ := domain.NewUser(email, username, password)

    if user == nil {
        t.Errorf("Error: %s", "User is nil")
    }

	repo := &mockUserRepository{
		users: make(map[string]*domain.User),
	}

    repo.InsertUser(user)

    LoginUseCase := NewLoginUseCase(repo)

    auth, err := LoginUseCase.Authenticate(username, password)

    if err != nil {
        t.Errorf("Error: %s", err)
    }

    if auth == false {
        t.Errorf("Error: %s", "Authenticate failed")
    }
}
