package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func (mcu *MockCommentUseCase) GetAll() ([]*domain.Comment, error) {
	comments := make([]*domain.Comment, 0)
	for _, comment := range mcu.Comments {
		comments = append(comments, comment)
	}
	return comments, nil
}

func (mcu *MockCommentUseCase) GetByID(commentID int64) (*domain.Comment, error) {
	comment, ok := mcu.Comments[commentID]
	if !ok {
		return nil, errors.New("GetByID: comment not found")
	} 
	return comment, nil
}

func (mcu *MockCommentUseCase) GetByPostID(postID int64) ([]*domain.Comment, error) {
	comments := make([]*domain.Comment, 0)
	for _, comment := range mcu.Comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments, nil
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

func newTestRouter(mcu *MockCommentUseCase) *chi.Mux {
	r := chi.NewRouter()
	commentHandler := NewCommentHandler(mcu)
	r.Route("/comments", func(r chi.Router) {
		r.Get("/", commentHandler.HandleGetAllComments)
		r.Get("/{commentID}", commentHandler.HandleGetCommentByID)
		r.Get("/post/{postID}", commentHandler.HandleGetCommentsByPostId)
		r.Post("/", commentHandler.HandleCreateComment)
		r.Put("/{commentID}", commentHandler.HandleUpdateComment)
		r.Delete("/{commentID}", commentHandler.HandleDeleteComment)
	})
	return r
}

func executeTestRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d", expected, actual)
	} else {
		println("Response code OK")
	}
}

func generateDummyComments(numComments int) map[int64]*domain.Comment {
	comments := make(map[int64]*domain.Comment)
	for i := int64(1); i <= int64(numComments); i++ {
		comments[i] = &domain.Comment{
			CommentID:     int64(i),
			UserID:        1, // o i si quieres que cada comentario tenga un usuario diferente
			PostID:        1, // o i si quieres que cada comentario tenga un post diferente
			Comment:       "This is a test comment " + strconv.Itoa(int(i)),
			InsertionDate: time.Now(),
			UpdateDate:    time.Now(),
		}
	}
	return comments
}

func TestCommentHandler_HandleGetAllComments(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("GET", "/comments", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCommentHandler_HandleGetCommentByID(t *testing.T) {
	mcu := NewMockCommentUseCase()
	mcu.Comments = generateDummyComments(2)
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("GET", "/comments/1", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCommentHandler_HandleGetCommentsByPostId(t *testing.T) {
	mcu := NewMockCommentUseCase()
	mcu.Comments = generateDummyComments(2)
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("GET", "/comments/post/1", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)	
}

func TestCommentHandler_HandleCreateComment(t *testing.T){
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
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
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

//CONTINUAR AQUÍ

func TestCommentHandler_HandleUpdateComment(t *testing.T) {
	mcu := NewMockCommentUseCase()
	mcu.Comments = generateDummyComments(2)
	r := newTestRouter(mcu)
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
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCommentHandler_HandleDeleteComment(t *testing.T) {
	mcu := NewMockCommentUseCase()
	mcu.Comments = generateDummyComments(2)
	r := newTestRouter(mcu)
	req, err := http.NewRequest("DELETE", "/comments/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCommentHandler_HandleGetAllComments_EmptyComments(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("GET", "/comments", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCommentHandler_HandleGetCommentByID_NotFound(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("GET", "/comments/1", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestCommentHandler_HandleCreateComment_InvalidData(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("POST", "/comments", bytes.NewReader([]byte("Invalid JSON")))
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestCommentHandler_HandleUpdateComment_InvalidData(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("PUT", "/comments/1", bytes.NewReader([]byte("Invalid JSON")))
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestCommentHandler_HandleDeleteComment_NotFound(t *testing.T) {
	mcu := NewMockCommentUseCase()
	r := newTestRouter(mcu)
	req, _ := http.NewRequest("DELETE", "/comments/1", nil)
	response := executeTestRequest(r, req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}