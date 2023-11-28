package infrastructure

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/auth/domain"
)

func TestTokenMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		mockValidateFn func(tokenString string) (*domain.User, error)
		wantStatus     int
		wantBody       string
	}{
		{
			name:       "No Authorization header",
			authHeader: "",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return nil, nil
			},
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"No token provided","status":"ERROR"}`,
		},
		{
			name:       "Invalid token",
			authHeader: "Bearer invalidToken",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return nil, application.ErrInvalidToken
			},
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"response":"Invalid token","status":"ERROR"}`,
		},
		{
			name:       "Valid token",
			authHeader: "Bearer validTokenString",
			mockValidateFn: func(tokenString string) (*domain.User, error) {
				return &domain.User{Username: "validUser", Role: "userRole"}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/some-protected-route", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			mockUseCase := &mockAuthUseCase{
				ValidateTokenFn: tt.mockValidateFn,
			}
			middleware := NewTokenMiddleware(mockUseCase)
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			testMiddleware := middleware.Middleware(nextHandler)
			recorder := httptest.NewRecorder()

			testMiddleware.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("%s: expected status %v; got %v", tt.name, tt.wantStatus, res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("%s: could not read response body: %v", tt.name, err)
			}

			if strings.TrimSpace(string(body)) != tt.wantBody {
				t.Errorf("%s: expected body %q; got %q", tt.name, tt.wantBody, body)
			}
		})
	}
}
