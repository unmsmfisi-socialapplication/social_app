package infraestructure

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_topics/application"
)

type mockInterestTopicUsecase struct {
	SetInterestTopicsFn func(userId string, interestId []string) error
}

func (m *mockInterestTopicUsecase) SetInterestTopics(userId string, interestId []string) error {
	return m.SetInterestTopicsFn(userId, interestId)
}
func TestHandleHandleSelectTopic(t *testing.T) {

	tests := []struct {
		name       string
		inputBody  string
		mock       func(userId string, interestId []string) error
		wantStatus int
		wantBody   string
	}{

		{
			name:       "Insertion Failed - user_id",
			inputBody:  `{"user_id":"123123","interest_id":["1","3","4"]}`,
			mock:       func(userId string, interestId []string) error { return errors.New("invalid insertion") },
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error during insertion","status":"ERROR"}`,
		},
		{
			name:       "Insertion Failed - interest_id",
			inputBody:  `{"user_id":"1","interest_id":["12","23","24"]}`,
			mock:       func(userId string, interestId []string) error { return errors.New("invalid insertion") },
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error during insertion","status":"ERROR"}`,
		},
		{
			name:       "Existing user interest topic",
			inputBody:  `{"user_id":"1","interest_id":["1","3","4"]}`,
			mock:       func(userId string, interestId []string) error { return application.ExistingUserInterestTopic },
			wantStatus: http.StatusConflict,
			wantBody:   `{"response":"Attempted insertion of an existing user interest topic","status":"ERROR"}`,
		},

		{
			name:       "Number of interest topics less than 3",
			inputBody:  `{"user_id":"1","interest_id":["1","3"]}`,
			mock:       func(userId string, interestId []string) error { return application.ExistingUserInterestTopic },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"At least 3 interest topics must be selected","status":"ERROR"}`,
		},

		{
			name:       "Duplicate interest_id",
			inputBody:  `{"user_id":"1","interest_id":["1","3","1"]}`,
			mock:       func(userId string, interestId []string) error { return application.ExistingUserInterestTopic },
			wantStatus: http.StatusConflict,
			wantBody:   `{"response":"Duplicate interest topic","status":"ERROR"}`,
		},
		{
			name:       "Bad request - General",
			inputBody:  `{"user_id": "1","interest_id": ""}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Bad request - user_id not defined",
			inputBody:  `{"interest_id": ["1","3","4"]}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Bad request - interest_id not defined",
			inputBody:  `{"user_id": "1"}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Bad request - extra parameters",
			inputBody:  `{"user_id":"1","interest_id":["1","3","4"],"random":""}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Insertion successful",
			inputBody:  `{"user_id":"1","interest_id":["1","3","4"]}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"Insertion successful","status":"OK"}`,
		},

		{
			name:       "Skipped insertion ",
			inputBody:  `{"user_id":"1","interest_id":[]}`,
			mock:       func(userId string, interestId []string) error { return nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"Skipped setting interest topics","status":"OK"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/interestTopics", bytes.NewBufferString(tt.inputBody))

			if req.URL.Path != "/interestTopics" {
				t.Errorf("expected request to /interestTopics, got %s", req.URL.Path)
				return
			}

			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockInterestTopicUsecase{
				SetInterestTopicsFn: tt.mock,
			}
			handler := NewSelectTopicHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleSelectTopic(recorder, req)

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
