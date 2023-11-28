package test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/infrastructure"
)

type mockPostUseCase struct {
	CreatePostFn func(post domain.PostCreate) (*domain.PostResponse, error)
	GetPostsFn   func(params domain.PostPaginationParams) (*domain.PostPagination, error)
	GetPostFn    func(id int) (*domain.Post, error)
	DeletePostFn func(id int64) error
}

func (m *mockPostUseCase) CreatePost(post domain.PostCreate) (*domain.PostResponse, error) {
	return m.CreatePostFn(post)
}

func (m *mockPostUseCase) GetPosts(params domain.PostPaginationParams) (*domain.PostPagination, error) {
	return m.GetPostsFn(params)
}

func (m *mockPostUseCase) GetPost(id int) (*domain.Post, error) {
	return m.GetPostFn(id)
}

func (m *mockPostUseCase) DeletePost(id int64) error {
	return m.DeletePostFn(id)
}

func TestHandleCreatePost(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockCreate func(post domain.PostCreate) (*domain.PostResponse, error)
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Invalid Request User",
			inputBody:  `{"title": "Test Post", "description": "Sample description", "hasMultimedia": false, "public": true, "multimedia": ""}`,
			mockCreate: func(post domain.PostCreate) (*domain.PostResponse, error) { return nil, nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request User","status":"ERROR"}`,
		},
		{
			name:       "Invalid Request Title",
			inputBody:  `{"userId": 1, "description": "Sample description", "hasMultimedia": false, "public": true, "multimedia": ""}`,
			mockCreate: func(post domain.PostCreate) (*domain.PostResponse, error) { return nil, nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request Title","status":"ERROR"}`,
		},
		{
			name:      "Internal Server Error",
			inputBody: `{"userId": 1, "title": "Test Post", "description": "Sample description", "hasMultimedia": false, "public": true, "multimedia": ""}`,
			mockCreate: func(post domain.PostCreate) (*domain.PostResponse, error) {
				return nil, fmt.Errorf("Internal Server Error")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Internal Server Error","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(tt.inputBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockPostUseCase{
				CreatePostFn: tt.mockCreate,
			}
			handler := infrastructure.NewPostHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleCreatePost(recorder, req)

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

func TestHandleGetAllPost(t *testing.T) {
	tests := []struct {
		name         string
		mockGetPosts func(params domain.PostPaginationParams) (*domain.PostPagination, error)
		wantStatus   int
		wantBody     string
	}{
		{
			name: "Internal Server Error",
			mockGetPosts: func(params domain.PostPaginationParams) (*domain.PostPagination, error) {
				return nil, fmt.Errorf("Internal Server Error")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Internal Server Error","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/getall", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockPostUseCase{
				GetPostsFn: tt.mockGetPosts,
			}
			handler := infrastructure.NewPostHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleGetAllPost(recorder, req)

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

// DELETE POST COMPONENT TESTS//
func TestHandleDeletePost(t *testing.T) {
	tests := []struct {
		name       string
		postID     int64
		mockDelete func(id int64) error
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Invalid Post ID",
			postID:     -2,
			mockDelete: func(id int64) error { return fmt.Errorf("Invalid post ID") },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid post ID","status":"ERROR"}`,
		},
		{
			name:       "Post Not Found",
			postID:     123,
			mockDelete: func(id int64) error { return fmt.Errorf("post not found") },
			wantStatus: http.StatusNotFound,
			wantBody:   `{"response":"Post not found","status":"ERROR"}`,
		},
		{
			name:       "Post Deleted Successfully",
			postID:     123,
			mockDelete: func(id int64) error { return nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"Post deleted successfully","status":"SUCCESS"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%v", tt.postID), nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", strconv.FormatInt(tt.postID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			mockUseCase := &mockPostUseCase{
				DeletePostFn: tt.mockDelete,
			}
			handler := infrastructure.NewPostHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleDeletePost(recorder, req)

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
