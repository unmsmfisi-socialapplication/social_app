package infrastructure

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

type mockAuthUseCase struct {
	application.AuthUseCaseInterface
	ValidateTokenFn func(tokenString string) (*domain.User, error)
	LogoutFn        func(username string) error
}

func (m *mockAuthUseCase) ValidateToken(tokenString string) (*domain.User, error) {
	return m.ValidateTokenFn(tokenString)
}

func (m *mockAuthUseCase) Logout(username string) error {
	return m.LogoutFn(username)
}

func TestAuthHandler_ValidateUserToken(t *testing.T) {
	tests := []struct {
		name         string
		tokenString  string
		mockValidate func(tokenString string) (*domain.User, error)
		wantStatus   int
		wantBody     string
	}{
		{
			name:        "No token provided",
			tokenString: "",
			mockValidate: func(tokenString string) (*domain.User, error) {
				return nil, nil
			},
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"No token provided","status":"ERROR"}`,
		},
		{
			name:        "Invalid token",
			tokenString: "invalidToken",
			mockValidate: func(tokenString string) (*domain.User, error) {
				return nil, application.ErrInvalidToken
			},
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"Invalid token","status":"ERROR"}`,
		},
		{
			name:        "Token validation error",
			tokenString: "validToken",
			mockValidate: func(tokenString string) (*domain.User, error) {
				return nil, errors.New("some validation error")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error during token validation","status":"ERROR"}`,
		},
		{
			name:        "Valid token",
			tokenString: "validToken",
			mockValidate: func(tokenString string) (*domain.User, error) {
				return &domain.User{Username: "mockUser", Role: "mockRole"}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":{"username":"mockUser","role":"mockRole"},"status":"OK"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/validate", nil)
			req.Header.Set("Authorization", tt.tokenString)

			mockUseCase := &mockAuthUseCase{
				ValidateTokenFn: tt.mockValidate,
			}
			handler := NewAuthHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.ValidateUserToken(recorder, req)

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
func TestAuthHandler_LogoutUser(t *testing.T) {
	tests := []struct {
		name           string
		tokenString    string
		mockValidateFn func(tokenString string) (*domain.User, error)
		mockLogoutFn   func(username string) error
		wantStatus     int
		wantBody       string
	}{
		{
			name:        "No token provided",
			tokenString: "",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return nil, nil
			},
			mockLogoutFn: nil,
			wantStatus:   http.StatusUnauthorized,
			wantBody:     `{"response":"No token provided","status":"ERROR"}`,
		},
		{
			name:        "Invalid token",
			tokenString: "Bearer invalidToken",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return nil, application.ErrInvalidToken
			},
			mockLogoutFn: nil,
			wantStatus:   http.StatusUnauthorized,
			wantBody:     `{"response":"Invalid token","status":"ERROR"}`,
		},
		{
			name:        "Error during token validation",
			tokenString: "Bearer validToken",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return nil, errors.New("validation error")
			},
			mockLogoutFn: nil,
			wantStatus:   http.StatusInternalServerError,
			wantBody:     `{"response":"Error during token validation","status":"ERROR"}`,
		},
		{
			name:        "Successful logout",
			tokenString: "Bearer validToken",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return &domain.User{Username: "validUser"}, nil
			},
			mockLogoutFn: func(username string) error {
				return nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"User logged out successfully","status":"OK"}`,
		},
		{
			name:        "Error during logout process",
			tokenString: "Bearer validToken",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return &domain.User{Username: "validUser"}, nil
			},
			mockLogoutFn: func(username string) error {
				return errors.New("logout error")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error during logout process","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/logout", nil)
			if tt.tokenString != "" {
				req.Header.Set("Authorization", tt.tokenString)
			}

			mockUseCase := &mockAuthUseCase{
				ValidateTokenFn: tt.mockValidateFn,
				LogoutFn:        tt.mockLogoutFn,
			}
			handler := NewAuthHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.LogoutUser(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("%s: expected status %v; got %v", tt.name, tt.wantStatus, res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("%s: could not read response body: %v", tt.name, err)
			}

			if strings.Trim(string(body), "\n") != tt.wantBody {
				t.Errorf("%s: expected body %q; got %q", tt.name, tt.wantBody, body)
			}
		})
	}
}
