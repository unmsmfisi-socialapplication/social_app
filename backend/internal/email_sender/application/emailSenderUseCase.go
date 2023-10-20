package application

import (
	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

type EmailSenderRepository interface {
	SendEmail(email *domain.Email) error
}

type EmailSenderUseCaseInterface interface {
	SendEmail(email *domain.Email) error
}
type EmailSenderUseCase struct {
	repository EmailSenderRepository
}

func NewEmailSenderUseCase(repo EmailSenderRepository) *EmailSenderUseCase {
	return &EmailSenderUseCase{
		repository: repo,
	}
}
func (es *EmailSenderUseCase) SendEmail(email *domain.Email) error {
	return es.repository.SendEmail(email)
}
