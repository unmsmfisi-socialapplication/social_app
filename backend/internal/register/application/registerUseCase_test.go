package application

import (
	"errors"
	"testing"

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

	useCase := NewRegistrationUseCase(repo)

	userCreate := domain.UserCreate{
		Email:    "user@example.com",
		Username: "user123",
		Password: "Password123!",
	}

	user, err := useCase.RegisterUser(userCreate)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if user == nil {
		t.Error("Expected a user, but got nil")
		return
	}
	if user.Email != userCreate.Email {
		t.Errorf("Expected user email to be %s, but got %s", userCreate.Email, user.Email)
	}

	_, err = useCase.RegisterUser(userCreate)
	if err != ErrEmailInUse {
		t.Errorf("Expected ErrEmailInUse, but got %v", err)
	}

	userCreate.Password = "password"
	userCreate.Email = "anotherEmai3.141516@mail.com"	
	
	_, err = useCase.RegisterUser(userCreate)
	if err != ErrFormat {
		t.Errorf("Expected ErrFormat, but got %v", err)
	}

}
