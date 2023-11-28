package application

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

type MockAuthSessionRepository struct {
	CheckValidSessionFunc func(username string) (bool, error)
	GetJTIByUsernameFunc  func(username string) (string, error)
	InvalidateSessionFunc func(username string) error
}

func (m *MockAuthSessionRepository) CheckValidSession(username string) (bool, error) {
	return m.CheckValidSessionFunc(username)
}

func (m *MockAuthSessionRepository) GetJTIByUsername(username string) (string, error) {
	return m.GetJTIByUsernameFunc(username)
}

func (m *MockAuthSessionRepository) InvalidateSession(username string) error {
	return m.InvalidateSessionFunc(username)
}

func MockParseToken(tokenString string) (*jwt.StandardClaims, error) {
	return &jwt.StandardClaims{
		Subject:   "mockUser",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        "mockedJTI",
	}, nil
}

func TestAuthUseCase_ValidateToken(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "secretkey")

	mockRepo := &MockAuthSessionRepository{
		GetJTIByUsernameFunc: func(username string) (string, error) {
			return "mockedJTI", nil
		},
	}

	authUseCase := NewAuthUseCase(mockRepo)

	tokenString := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJtb2NrVXNlciIsImp0aSI6Im1vY2tlZEpUSSIsImV4cCI6MTcwMTU2ODgwMH0.6iwPkwamyoF8MmD3oKqiFlMPcmb-8HTcD5uPwnsBF7E"

	user, err := authUseCase.ValidateToken(tokenString)

	expectedUsername := "mockUser"
	if err != nil {
		t.Errorf("ValidateToken() error = %v, wantErr %v", err, nil)
	}
	if user == nil || user.Username != expectedUsername {
		t.Errorf("Expected user to be '%v', got %v", expectedUsername, user)
	}
}
