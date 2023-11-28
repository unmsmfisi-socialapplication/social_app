package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func mockAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") == "valid_token" {
			ctx := r.Context()
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"success"}`))
}

func mockRouter() http.Handler {
	freeRoutes := chi.NewRouter()

	protectedRoutes := chi.NewRouter()
	protectedRoutes.Use(mockAuthMiddleware)
	protectedRoutes.Get("/protected", mockHandler)

	freeRoutes.Get("/free", mockHandler)
	freeRoutes.Mount("/", protectedRoutes)

	return freeRoutes
}

func TestProtectedRoutes(t *testing.T) {
	r := mockRouter()

	tests := []struct {
		name       string
		route      string
		method     string
		headers    map[string]string
		wantStatus int
	}{
		{
			name:       "Protected Route with Authentication header (valid token)",
			route:      "/protected",
			method:     http.MethodGet,
			headers:    map[string]string{"Authorization": "valid_token"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "Protected Route without Authentication",
			route:      "/protected",
			method:     http.MethodGet,
			headers:    nil,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Protected Route with Authentication header (invalid token)",
			route:      "/protected",
			method:     http.MethodGet,
			headers:    map[string]string{"Authorization": "invalid_token"},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Free Route",
			route:      "/free",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			for key, value := range tc.headers {
				req.Header.Set(key, value)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if rr.Code != tc.wantStatus {
				t.Errorf("%s: expected status code %v, got %v", tc.name, tc.wantStatus, rr.Code)
			}
		})
	}
}
