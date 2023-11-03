package domain

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name    string
		to      string
		subject string
		body    string
		wantErr bool
	}{
		{
			name:    "success",
			to:      "test@testappsocial.com",
			subject: "Test Subject",
			body:    "Test Body",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.to, tt.subject, tt.body)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && (email.To != tt.to || email.Subject != tt.subject || email.Body != tt.body) {
				t.Errorf("NewEmail() got = %v, want to=%v, subject=%v, body=%v", email, tt.to, tt.subject, tt.body)
			}
		})
	}
}
