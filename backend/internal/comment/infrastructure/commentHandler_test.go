package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/unmsmfisi-socialapplication/social_app/internal/comment/domain"
)

// To simulate CommentUseCaseInterface
type MockCommentUseCase struct {
	Comments map[int64]*domain.Comment
}

func NewMockCommentUseCase() *MockCommentUseCase {
	return &MockCommentUseCase{
		Comments: make(map[int64]*domain.Comment),
	}
}

func (mcu *MockCommentUseCase) GetByID(commentID int64) (*domain.Comment, error) {
	comment, ok := mcu.Comments[commentID]
	if !ok {
		return nil, errors.New("GetByID: comment not found")
	} 
	return comment, nil
}

func (mcu *MockCommentUseCase) Create(comment *domain.Comment) error {
	mcu.Comments[comment.CommentID] = comment
	return nil
}

func (mcu *MockCommentUseCase) Update(commentID int64, comment *domain.Comment) error {
	_, ok := mcu.Comments[commentID]
	if !ok {
		return errors.New("Update: Comment not found")
	}

	mcu.Comments[commentID] = comment
	return nil
}

func (mcu *MockCommentUseCase) Delete(commentID int64) error {
	_, ok := mcu.Comments[commentID]
	if !ok {
		return errors.New("Delete: Comment not found")
	}

	delete(mcu.Comments, commentID)
	return nil
}

func setupRouter() *chi.Mux {
    r := chi.NewRouter()
    return r
}

func TestCommentHandler_HandleGetCommentByID(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()
    	
	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "Este es un comentario de prueba",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}
	commentHandler := NewCommentHandler(mockUseCase)
	r.Get("/comments/{commentID}", commentHandler.HandleGetCommentByID)
	
	req, err := http.NewRequest("GET", "/comments/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusOK, w.Code)
	} else {
		println("Comment retrieved successfully")
	}
}

func TestCommentHandler_HandleCreateComment(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()
	commentHandler := NewCommentHandler(mockUseCase)

	r.Post("/comments", commentHandler.HandleCreateComment)

	commentData := struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`
	}{
		UserID:          1,
		PostID:          2,
		Comment:         "Test Comment",
		InsertionDate:   time.Now(),
		UpdateDate:      time.Now(),
		ParentCommentID: 0,
	}
	requestBody, _ := json.Marshal(commentData)
	req, err := http.NewRequest("POST", "/comments", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusOK, w.Code)
	} else {
		println("Comment created successfully")
	}
}

func TestCommentHandler_HandleUpdateComment(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()

	// Add simulate comments to map 
	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "This is a simple comment",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	commentHandler := NewCommentHandler(mockUseCase)
	r.Put("/comments/{commentID}", commentHandler.HandleUpdateComment)
	commentData := struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`
	}{
		UserID:          1,
		PostID:          2,
		Comment:         "Updated Comment",
		InsertionDate:   time.Now(),
		UpdateDate:      time.Now(),
		ParentCommentID: 0,
	}
	requestBody, _ := json.Marshal(commentData)
	req, err := http.NewRequest("PUT", "/comments/1", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusOK, w.Code)
	} else {
		println("Comment updated successfully")
	}
}

func TestCommentHandler_HandleDeleteComment(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()

	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "Este es un comentario de prueba",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	commentHandler := NewCommentHandler(mockUseCase)
	r.Delete("/comments/{commentID}", commentHandler.HandleDeleteComment)

	req, err := http.NewRequest("DELETE", "/comments/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusOK, w.Code)
	} else {
		println("Comment deleted successfully")
	}
}

func TestCommentHandler_HandleGetCommentByID_NotFound(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()

	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "This is a simple comment",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	commentHandler := NewCommentHandler(mockUseCase)
	r.Get("/comments/{commentID}", commentHandler.HandleGetCommentByID)
	
	req, err := http.NewRequest("GET", "/comments/99", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusNotFound, w.Code)
	} else {
		println("Comment not found as expected")
	}
}

func TestCommentHandler_HandleCreateComment_InvalidData(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()
	commentHandler := NewCommentHandler(mockUseCase)

	r.Post("/comments", commentHandler.HandleCreateComment)

	//The invalid comment data is missing the comment field
	invalidCommentData := struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         int64     `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`
	}{
		UserID:          1,
		PostID:          2,
		Comment:         100,
		InsertionDate:   time.Now(),
		UpdateDate:      time.Now(),
		ParentCommentID: 0,
	}

	requestBody, _ := json.Marshal(invalidCommentData)
	req, err := http.NewRequest("POST", "/comments", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusBadRequest, w.Code)
	} else {
		println("Invalid comment data as expected")
	}
}

func TestCommentHandler_HandleUpdateComment_NotFound(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()

	// Add simulate comments to map 
	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "This is a simple comment",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	commentHandler := NewCommentHandler(mockUseCase)

	r.Put("/comments/{commentID}", commentHandler.HandleUpdateComment)

	commentData := struct {
		UserID          int64     `json:"userID"`
		PostID          int64     `json:"postID"`
		Comment         string    `json:"comment"`
		InsertionDate   time.Time `json:"insertionDate"`
		UpdateDate      time.Time `json:"updateDate"`
		ParentCommentID int64     `json:"parentCommentID"`
	}{
		UserID:          1,
		PostID:          2,
		Comment:         "Updated Comment",
		InsertionDate:   time.Now(),
		UpdateDate:      time.Now(),
		ParentCommentID: 0,
	}

	requestBody, _ := json.Marshal(commentData)
	req, err := http.NewRequest("PUT", "/comments/9999", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusNotFound, w.Code)
	} else {
		println("Comment not found for update as expected")
	}
}

func TestCommentHandler_HandleDeleteComment_NotFound(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()

	commentID := int64(1)
	mockUseCase.Comments[commentID] = &domain.Comment{
		CommentID:  commentID,
		UserID:     1,
		PostID:     1,
		Comment:    "Este es un comentario de prueba",
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}

	commentHandler := NewCommentHandler(mockUseCase)
	r.Delete("/comments/{commentID}", commentHandler.HandleDeleteComment)

	req, err := http.NewRequest("DELETE", "/comments/9999", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusNotFound, w.Code)
	} else {
		println("Comment not found for delete as expected")
	}
}

// Future test when start integration for integration with POSTS
/*
func TestCommentHandler_HandleGetCommentsByPostID(t *testing.T) {
	r := chi.NewRouter()
	mockUseCase := NewMockCommentUseCase()
	commentHandler := NewCommentHandler(mockUseCase)
	r.Get("/posts/{postID}/comments", commentHandler.HandleGetCommentsByPostID)

	postID := int64(1)
	req, err := http.NewRequest("GET", "/posts/1/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ERROR: Waiting status code %d, but get %d", http.StatusOK, w.Code)
	}
}*/