package security

import "net/mail"

type EmailService struct {
}

func NewEmailService() EmailService {
	return EmailService{}
}

func (service EmailService) IsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}