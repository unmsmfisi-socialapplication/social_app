package infrastructure

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/application"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

type mockEmailSenderUsecase struct {
	application.EmailSenderUseCase
	SendEmailFn func(email *domain.Email) error
}

func (m *mockEmailSenderUsecase) SendEmail(email *domain.Email) error {
	return m.SendEmailFn(email)
}

func TestHandleSendEmail(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockSend   func(email *domain.Email) error
		wantStatus int
		wantBody   string
	}{

		{
			name:       "ERROR - Invalid request payload (Sending different fields to expected on JSON request)",
			inputBody:  `{"To": "test@testappsocial.com", "Title": "test", "Body": "test message"}`,
			mockSend:   func(email *domain.Email) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "ERROR - Invalid request payload ",
			inputBody:  `{"To": "test@testappsocial.com", "Subject": "test", "Body": "test message" `,
			mockSend:   func(email *domain.Email) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "ERROR - Invalid request payload (Sending only two fields  on JSON request) ",
			inputBody:  `{"To": "test@testappsocial.com", "Body": "test message"} `,
			mockSend:   func(email *domain.Email) error { return nil },
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"response":"Invalid request payload","status":"ERROR"}`,
		},

		{
			name:       "Internal Server Error",
			inputBody:  `{"To": "test@testappsocial.com", "Subject": "test", "Body": "test message"}`,
			mockSend:   func(email *domain.Email) error { return errors.New("ERROR 500") },
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"response":"Internal server error","status":"ERROR"}`,
		},
		{
			name:       "Email Sent Successfully",
			inputBody:  `{"To": "test@testappsocial.com", "Subject": "test", "Body": "test message"}`,
			mockSend:   func(email *domain.Email) error { return nil },
			wantStatus: http.StatusOK,
			wantBody:   `{"response":"Email sent successfully","status":"OK"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/send", bytes.NewBufferString(tt.inputBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			mockUseCase := &mockEmailSenderUsecase{
				SendEmailFn: tt.mockSend,
			}
			handler := NewEmailSenderHandler(mockUseCase)
			recorder := httptest.NewRecorder()

			handler.HandleSendEmail(recorder, req)

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
