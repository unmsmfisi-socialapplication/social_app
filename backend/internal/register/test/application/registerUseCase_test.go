package applicationtest

import (
	"errors"
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

func TestRegistrationUseCase_RegisterUser(t *testing.T) {
	
	repo := &mockUserRepository{
		users: make(map[string]*domain.User),
	}

	useCase := application.NewRegistrationUseCase(repo)

	
	phone := "123456789"
	email := "user@example.com"
	username := "user123"
	password := "Password123!"

	user, err := useCase.RegisterUser(phone, email, username, password)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if user == nil {
		t.Error("Expected a user, but got nil")
		return
	}
	if user.Email != email {
		t.Errorf("Expected user email to be %s, but got %s", email, user.Email)
	}

	
	_, err = useCase.RegisterUser(phone, email, username, password)
	if err != application.ErrEmailInUse {
		t.Errorf("Expected ErrEmailInUse, but got %v", err)
	}

	
	invalidPassword := "password"
	_, err = useCase.RegisterUser(phone, "new@example.com", "newuser", invalidPassword)
	if err != application.ErrFormat {
		t.Errorf("Expected ErrFormat, but got %v", err)
	}

	
	_, err = useCase.RegisterUser("1234567890", "new@example.com", "newuser", "Valid123!")
	if err != application.ErrPhone {
		t.Errorf("Expected ErrPhone, but got %v", err)
	}
}
