package infrastructure_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	login_app "github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
	login_dom "github.com/unmsmfisi-socialapplication/social_app/internal/login/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/infrastructure"
)

type UserRepositoryMock struct {
	users map[string]*domain.User    // Simulate a memory database
	sessions       map[string]string // Map to simulate sessions by username
	existingAccess map[string]bool   // Map to simulate existing access by username
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		users:          make(map[string]*domain.User),
		sessions:       make(map[string]string),
		existingAccess: make(map[string]bool),
	}
}

func (ur *UserRepositoryMock) GetUserByEmail(email string) (*domain.User, error) {
	user, ok := ur.users[email]
	if !ok {
		return nil, application.ErrUserNotFound
	}
	return user, nil
}

func (ur *UserRepositoryMock) InsertUser(newUser *domain.User) (*domain.User, error) {
	if _, ok := ur.users[newUser.Email]; ok {
		return nil, application.ErrEmailInUse
	}

	ur.users[newUser.Email] = newUser
	return newUser, nil
}

func (ur *UserRepositoryMock) GetUserByUsername(username string) (*login_dom.User, error) {
    return nil, nil
}

func (ur *UserRepositoryMock) StoreJTIForSession(username string, jti string) error {
	ur.sessions[username] = jti
	return nil
}

func (ur *UserRepositoryMock) InsertSession(username string) error {
	ur.existingAccess[username] = true
	return nil
}

func (ur *UserRepositoryMock) UpdateSession(username string) error {
	return nil
}

func (ur *UserRepositoryMock) CheckExistingSession(username string) (bool, error) {
	_, exists := ur.existingAccess[username]
	return exists, nil
}


func TestRegisterUserHandler_RegisterUser(t *testing.T) {
    os.Setenv("JWT_SECRET_KEY", "test")
    mockRepository := NewUserRepositoryMock()
    useCase := application.NewRegistrationUseCase(mockRepository)
    loginUseCase := login_app.NewLoginUseCase(mockRepository)
	handler := infrastructure.NewRegisterUserHandler(useCase, loginUseCase)

	r := http.NewServeMux()
	r.HandleFunc("/register", handler.RegisterUser)

	t.Run("Successful registration", func(t *testing.T) {
		requestBody := map[string]string{
			"email":      "test@example.com",
			"user_name":  "testuser",
			"password":   "MyP@ssw0rd!",
		}
		reqBodyBytes, _ := json.Marshal(requestBody)

		req := httptest.NewRequest("POST", "/register", bytes.NewReader(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		// Check the status code
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d; got %d", http.StatusOK, rec.Code)
		}
        
        // Check the response body is what we expect.
        if !bytes.Contains(rec.Body.Bytes(), []byte("email")) {
            t.Errorf("expected email field")
        }
        if !bytes.Contains(rec.Body.Bytes(), []byte("user_name")) {
            t.Errorf("expected user_name field")
        }
        if !bytes.Contains(rec.Body.Bytes(), []byte("token_result")) {
            t.Errorf("expected token_result field")
        }

	})
}

