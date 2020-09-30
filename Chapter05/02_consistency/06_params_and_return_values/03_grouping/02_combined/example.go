package _2_combined

import (
	"net/smtp"
	"os"
)

type Sender struct {
	Server string
	Auth   smtp.Auth
}

func (s Sender) Send(from string, to string, email Email) error {
	// implementation removed
	return nil
}

type Email struct {
	Subject     string
	Message     string
	Attachments []os.File
}
