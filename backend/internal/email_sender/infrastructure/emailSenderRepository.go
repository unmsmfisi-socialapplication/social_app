package infrastructure

import (
	"errors"
	"net/smtp"

	"github.com/unmsmfisi-socialapplication/social_app/internal/email_sender/domain"
)

type loginAuth struct {
	username, password string
}

type EmailSenderRepository struct {
	auth   smtp.Auth
	sender EmailSender
}
type EmailSender interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}
type RealEmailSender struct{}

func (res *RealEmailSender) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func NewEmailSenderRepository(auth smtp.Auth, sender EmailSender) *EmailSenderRepository {
	return &EmailSenderRepository{
		auth:   auth,
		sender: sender,
	}
}

func (esr *EmailSenderRepository) SendEmail(email *domain.Email) error {
	msg := []byte("To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		email.Body + "\r\n")
	return esr.sender.SendMail("smtp.gmail.com:587", esr.auth, "contact.appsocial@gmail.com", []string{email.To}, msg)
}
