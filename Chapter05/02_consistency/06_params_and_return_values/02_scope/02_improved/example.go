package _2_improved

import (
	"net/smtp"
	"os"
)

type Sender struct {
	Server string
	Auth   smtp.Auth
}

func (s Sender) Send(from string, to string, subject string, message string, attachments []os.File) error {
	// implementation removed
	return nil
}
