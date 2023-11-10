package infrastructure

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/domain"
)

type mockListInterestTopicsUsecase struct {
	GetInteresTopicsFn func() ([]domain.InterestTopic, error)
}

func (m *mockListInterestTopicsUsecase) GetInteresTopics() ([]domain.InterestTopic, error) {
	return m.GetInteresTopicsFn()
}

func TestHandleListTopics(t *testing.T) {
	tests := []struct {
		name       string
		mock       func() ([]domain.InterestTopic, error)
		wantStatus int
		wantBody   string
	}{
		{
			name: "Data Found",
			mock: func() ([]domain.InterestTopic, error) {
				return []domain.InterestTopic{
					{InterestId: "1", InterestName: "Danza", InterestSummary: "Danza"},
					{InterestId: "2", InterestName: "Deportes", InterestSummary: "Deportes"},
					{InterestId: "3", InterestName: "Programacion", InterestSummary: "Programacion"},
				}, nil
			},
			wantStatus: http.StatusOK,
			wantBody: `{"response":[{"interest_id":"1","interest_name":"Danza","interest_summary":"Danza"},{"interest_id":"2","interest_name":"Deportes","interest_summary":"Deportes"},{"interest_id":"3","interest_name":"Programacion","interest_summary":"Programacion"}],"status":"OK"}`,
		},
		{
			name: "No Data Found",
			mock: func() ([]domain.InterestTopic, error) {
				return nil, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"There are no interest topics","status":"OK"}`,
		},
		{
			name: "Error with Database",
			mock: func() ([]domain.InterestTopic, error) {
				return nil, errors.New("Error while fetching data")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error while fetching data","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mockListInterestTopicsUsecase{
				GetInteresTopicsFn: tt.mock,
			}
			handler := NewListInterestTopicsHandler(mockUseCase)
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/interestTopics/list", nil)

			if req.URL.Path != "/interestTopics/list" {
				t.Errorf("expected request to /interestTopics/list, got %s", req.URL.Path)
				return
			}

			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			handler.HandleListTopics(recorder, req)

			res := recorder.Result()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("expected status %v; got %v", tt.wantStatus, res.StatusCode)
			}

			body := recorder.Body.String()

			if body != tt.wantBody {
				t.Errorf("expected body %q; got %q", tt.wantBody, body)
			}
		})
	}
}
