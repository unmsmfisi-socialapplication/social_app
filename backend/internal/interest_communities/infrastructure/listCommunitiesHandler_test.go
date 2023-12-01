package infrastructure

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/interest_communities/domain"
)

type mockListCommunitiesUseCase struct {
	GetCommunitiesListFn func(userId, pageSize, pageNumber string) ([]domain.Community, error)
}

func (m *mockListCommunitiesUseCase) GetCommunitiesList(userId, pageSize, pageNumber string) ([]domain.Community, error) {
	return m.GetCommunitiesListFn(userId, pageSize, pageNumber)
}

func TestHandleListCommunities(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mock       func(userId, pageSize, pageNumber string) ([]domain.Community, error)
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Bad request - Empty parameters",
			inputBody: `{"user_id":"","page_number":"","page_size":""}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Bad request - user id not defined",
			inputBody: `{"page_number":"","page_size":""}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Bad request - page number not defined",
			inputBody: `{"user_id":"","page_size":""}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Bad request - page number not defined",
			inputBody: `{"user_id":"","page_number":""}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Bad request - Empty parameters",
			inputBody: `{"user_id":"","page_number":"","page_size":""}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Bad request - Extra parameters",
			inputBody: `{"user_id":"1","page_number":"1","page_size":"2","extra":"extra"}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},
		{
			name:      "Data Found - Page 1",
			inputBody: `{"user_id":"1","page_number":"1","page_size":"2"}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return []domain.Community{
					{CommunityId: "1", CommunityName: "Programmers", CommunityDescription: "Programmers community", InterestId: "1"},
					{CommunityId: "2", CommunityName: "Musicians", CommunityDescription: "Musicians community", InterestId: "2"},
				}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":[{"community_id":"1","community_name":"Programmers","community_description":"Programmers community","interest_id":"1"},{"community_id":"2","community_name":"Musicians","community_description":"Musicians community","interest_id":"2"}],"status":"OK"}`,
		},
		{
			name:      "Data Found - Page 2",
			inputBody: `{"user_id":"1","page_number":"2","page_size":"2"}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return []domain.Community{
					{CommunityId: "3", CommunityName: "Engineers", CommunityDescription: "Engineers community", InterestId: "3"},
					{CommunityId: "4", CommunityName: "Singers", CommunityDescription: "Singers community", InterestId: "2"},
				}, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":[{"community_id":"3","community_name":"Engineers","community_description":"Engineers community","interest_id":"3"},{"community_id":"4","community_name":"Singers","community_description":"Singers community","interest_id":"2"}],"status":"OK"}`,
		},
		{
			name:      "No Data Found",
			inputBody: `{"user_id":"1","page_number":"2","page_size":"10"}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, nil
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"There are no communities","status":"OK"}`,
		},
		{
			name:      "Error with Database",
			inputBody: `{"user_id":"1","page_number":"2","page_size":"2"}`,
			mock: func(userId, pageSize, pageNumber string) ([]domain.Community, error) {
				return nil, errors.New("Error while fetching data")
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Error while fetching data","status":"ERROR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/interestCommunities/list", bytes.NewBufferString(tt.inputBody))

			if req.URL.Path != "/interestCommunities/list" {
				t.Errorf("expected request to /interestCommunities/list, got %s", req.URL.Path)
				return
			}

			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockListCommunitiesUseCase{
				GetCommunitiesListFn: tt.mock,
			}
			handler:= NewListCommunitiesHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleListCommunities(recorder, req)

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