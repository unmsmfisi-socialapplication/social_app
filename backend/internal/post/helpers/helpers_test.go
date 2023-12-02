package helpers

import (
	"net/url"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

func TestParsePaginationParams(t *testing.T) {
	testCases := []struct {
		name           string
		query          url.Values
		expectedParams domain.PostPaginationParams
		expectedError  bool
	}{
		{
			name: "Valid Query Parameters",
			query: url.Values{
				"page":  []string{"1"},
				"limit": []string{"10"},
			},
			expectedParams: domain.PostPaginationParams{
				Page:  1,
				Limit: 10,
			},
			expectedError: false,
		},
		{
			name: "Missing Page Parameter",
			query: url.Values{
				"limit": []string{"10"},
			},
			expectedParams: domain.PostPaginationParams{
				Page:  0,
				Limit: 10,
			},
			expectedError: false,
		},
		{
			name: "Invalid Page Parameter",
			query: url.Values{
				"page":  []string{"invalid"},
				"limit": []string{"10"},
			},
			expectedParams: domain.PostPaginationParams{
				Page:  0,
				Limit: 0,
			},
			expectedError: true,
		},
		{
			name: "Missing Limit Parameter",
			query: url.Values{
				"page": []string{"1"},
			},
			expectedParams: domain.PostPaginationParams{
				Page:  1,
				Limit: 0,
			},
			expectedError: false,
		},
		{
			name: "Invalid Limit Parameter",
			query: url.Values{
				"page":  []string{"1"},
				"limit": []string{"invalid"},
			},
			expectedParams: domain.PostPaginationParams{
				Page:  1,
				Limit: 0,
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := ParsePaginationParams(tc.query)

			if tc.expectedError && err == nil {
				t.Error("Expected an error, but got none")
			} else if !tc.expectedError && err != nil {
				t.Errorf("Expected no error, but got one: %v", err)
			}

			if tc.expectedParams != params {
				t.Errorf("Expected parsed pagination parameters %+v, but got %+v", tc.expectedParams, params)
			}
		})
	}
}
