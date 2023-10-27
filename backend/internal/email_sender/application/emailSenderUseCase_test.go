package application

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

type mockRepository struct {
	sendEmailErr  error
	emailReceived *domain.Email
}

func (m *mockRepository) SendEmail(email *domain.Email) error {
	m.emailReceived = email
	return m.sendEmailErr
}

func TestEmailSenderUseCase_SendEmail(t *testing.T) {
	email := &domain.Email{
		To:      "test@testappsocial.com",
		Subject: "Test",
		Body:    "This is a test email",
	}

	t.Run("delegates to repository", func(t *testing.T) {
		mockRepo := &mockRepository{}
		useCase := NewEmailSenderUseCase(mockRepo)

		_ = useCase.SendEmail(email)

		if mockRepo.emailReceived != email {
			t.Error("SendEmail did not delegate to the repository correctly")
		}
	})

	t.Run("handles repository error", func(t *testing.T) {
		mockError := errors.New("test error")
		mockRepo := &mockRepository{sendEmailErr: mockError}
		useCase := NewEmailSenderUseCase(mockRepo)

		err := useCase.SendEmail(email)

		if err != mockError {
			t.Errorf("expected error %v, got %v", mockError, err)
		}
	})
}
