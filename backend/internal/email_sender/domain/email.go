package domain

type Email struct {
	To      string
	Subject string
	Body    string
}

func NewEmail(to, subject, body string) (*Email, error) {
	return &Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}, nil
}
