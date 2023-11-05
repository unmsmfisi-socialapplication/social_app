package infrastructure

import (
	"errors"
	"net/smtp"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

type MockEmailSender struct {
	SendMailFn func(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

func (m *MockEmailSender) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return m.SendMailFn(addr, a, from, to, msg)
}

func TestSendEmail(t *testing.T) {
	mockAuth := LoginAuth("testuser", "testpass")

	tests := []struct {
		name    string
		sendErr error
		wantErr error
	}{
		{
			name:    "success",
			sendErr: nil,
			wantErr: nil,
		},
		{
			name:    "failure",
			sendErr: errors.New("failed sending email"),
			wantErr: errors.New("failed sending email"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSender := &MockEmailSender{
				SendMailFn: func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
					return tt.sendErr
				},
			}
			repo := NewEmailSenderRepository(mockAuth, mockSender)

			err := repo.SendEmail(&domain.Email{
				To:      "test@example.com",
				Subject: "Test Subject",
				Body:    "Test Body",
			})

			if err != tt.wantErr && err.Error() != tt.wantErr.Error() {
				t.Errorf("wanted error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
