package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/application"

	"github.com/unmsmfisi-socialapplication/social_app/internal/login/infrastructure"
)

type mockLoginUsecase struct {
	AuthenticateFn func(username, password string) (bool, error)
}

func (m *mockLoginUsecase) Authenticate(username, password string) (bool, error) {
	return m.AuthenticateFn(username, password)
}

func TestHandleLogin(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockAuth   func(username, password string) (bool, error)
		wantStatus int
		wantBody   string
	}{

		{
			name:       "User not found",
			inputBody:  `{"username": "nonexistent", "password": "password"}`,
			mockAuth:   func(username, password string) (bool, error) { return false, application.ErrUserNotFound },
			wantStatus: http.StatusNotFound,
			wantBody:   `{"response":"User not found","status":"NOTFOUND"}`,
		},
		{
			name:       "Invalid password",
			inputBody:  `{"username": "test", "password": "wrongpassword"}`,
			mockAuth:   func(username, password string) (bool, error) { return false, application.ErrInvalidCredentials },
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"Invalid password","status":"NOPASSWORD"}`,
		},
		{
			name:       "Bad request",
			inputBody:  `{"username": "test"`,
			mockAuth:   func(username, password string) (bool, error) { return false, nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:       "Authentication Failed",
			inputBody:  `{"username":"validuser", "password":"validformatpassword"}`,
			mockAuth:   func(username, password string) (bool, error) { return false, nil },
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"Authentication failed","status":"ERROR"}`,
		},

		{
			name:       "Login successful",
			inputBody:  `{"username": "test", "password": "test"}`,
			mockAuth:   func(username, password string) (bool, error) { return true, nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"Authentication successful","status":"OK"}`,
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
			}
			handler := infrastructure.NewLoginHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleLogin(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("expected status %v; got %v", tt.wantStatus, res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response body: %v", err)
			}

			if string(body) != tt.wantBody {
				t.Errorf("expected body %q; got %q", tt.wantBody, body)
			}
		})
	}
}
