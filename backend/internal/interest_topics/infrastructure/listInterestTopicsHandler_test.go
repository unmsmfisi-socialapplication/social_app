package infrastructure

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type mockListInterestTopicsUsecase struct {
	GetInteresTopicsFn func(pageSize, pageNumber string) ([]domain.InterestTopic, error)
}

func (m *mockListInterestTopicsUsecase) GetInteresTopics(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
	return m.GetInteresTopicsFn(pageSize, pageNumber)
}

func TestHandleListTopics(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mock       func(pageSize, pageNumber string) ([]domain.InterestTopic, error)
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Data Found - Page 1",
			inputBody: `{"page_number":"1","page_size":"2"}`,
			mock: func(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
				return []domain.InterestTopic{
					{InterestId: "1", InterestName: "Danza", InterestSummary: "Danza"},
					{InterestId: "2", InterestName: "Deportes", InterestSummary: "Deportes"},
				}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":[{"interest_id":"1","interest_name":"Danza","interest_summary":"Danza"},{"interest_id":"2","interest_name":"Deportes","interest_summary":"Deportes"}],"status":"OK"}`,
		},
		{
			name:      "Data Found- Page 2",
			inputBody: `{"page_number":"2","page_size":"2"}`,
			mock: func(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
				return []domain.InterestTopic{
					{InterestId: "3", InterestName: "Musica", InterestSummary: "Musica"},
					{InterestId: "4", InterestName: "Programacion", InterestSummary: "Programacion"},
				}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":[{"interest_id":"3","interest_name":"Musica","interest_summary":"Musica"},{"interest_id":"4","interest_name":"Programacion","interest_summary":"Programacion"}],"status":"OK"}`,
		},
		{
			name:      "No Data Found",
			inputBody: `{"page_number":"2","page_size":"10"}`,
			mock: func(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
				return nil, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"There are no interest topics","status":"OK"}`,
		},
		{
			name:      "Error with Database",
			inputBody: `{"page_number":"2","page_size":"2"}`,
			mock: func(pageSize, pageNumber string) ([]domain.InterestTopic, error) {
				return nil, errors.New("Error while fetching data")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error while fetching data","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/interestTopics/list", bytes.NewBufferString(tt.inputBody))

			if req.URL.Path != "/interestTopics/list" {
				t.Errorf("expected request to /interestTopics/list, got %s", req.URL.Path)
				return
			}

			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockListInterestTopicsUsecase{
				GetInteresTopicsFn: tt.mock,
			}
			handler := NewListInterestTopicsHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleListTopics(recorder, req)

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
