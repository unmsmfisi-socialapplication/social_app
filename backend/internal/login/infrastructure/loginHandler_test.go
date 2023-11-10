package infrastructure

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"
)

type mockLoginUsecase struct {
	AuthenticateFn       func(username, password string) (string, error)
	ExtractJTIFn         func(token string) (string, error)
	StoreJTIForSessionFn func(username string, jti string) error
}

func (m *mockLoginUsecase) Authenticate(username, password string) (string, error) {
	return m.AuthenticateFn(username, password)
}
func (m *mockLoginUsecase) ExtractJTI(token string) (string, error) {
	return m.ExtractJTIFn(token)
}

func (m *mockLoginUsecase) StoreJTIForSession(username string, jti string) error {
	return nil
}
func TestHandleLogin(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockAuth   func(username, password string) (string, error)
		wantStatus int
		wantBody   string
	}{
		{
			name:       "User not found",
			inputBody:  `{"username": "nonexistent", "password": "password"}`,
			mockAuth:   func(username, password string) (string, error) { return "", application.ErrUserNotFound },
			wantStatus: http.StatusNotFound,
			wantBody:   `{"response":"User not found","status":"NOTFOUND"}`,
		},
		{
			name:       "Invalid password",
			inputBody:  `{"username": "test", "password": "wrongpassword"}`,
			mockAuth:   func(username, password string) (string, error) { return "", application.ErrInvalidCredentials },
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"Invalid credentials","status":"BADCREDENTIALS"}`,
		},
		{
			name:       "Bad request",
			inputBody:  `{"username": "test"`,
			mockAuth:   func(username, password string) (string, error) { return "", nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Bad request",
			inputBody:  `{"username": "test","hashedpassword:"test"}`,
			mockAuth:   func(username, password string) (string, error) { return "", nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:       "Authentication Failed",
			inputBody:  `{"username":"validuser", "password":"validformatpassword"}`,
			mockAuth:   func(username, password string) (string, error) { return "", errors.New("Internal error") },
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error during authentication","status":"ERROR"}`,
		},
		{
			name:       "Login successful",
			inputBody:  `{"username": "myuser12", "password": "Social@12"}`,
			mockAuth:   func(username, password string) (string, error) { return "tokenString", nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":{"token":"tokenString"},"status":"OK"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(tt.inputBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockLoginUsecase{
				AuthenticateFn: tt.mockAuth,
				ExtractJTIFn: func(token string) (string, error) {
					return "someJTI", nil
				},
			}
			handler := NewLoginHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleLogin(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("expected status %v; got %v", tt.wantStatus, res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if string(body) != tt.wantBody {
				t.Errorf("expected body %q; got %q", tt.wantBody, body)
			}
		})
	}
}
