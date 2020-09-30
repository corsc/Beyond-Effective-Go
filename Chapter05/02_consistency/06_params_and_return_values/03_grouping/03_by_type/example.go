package _2_improved

import (
	"net/smtp"
	"os"
)

type Sender struct {
	Server string
	Auth   smtp.Auth
}

func (s Sender) Send(from, to string, message Message) error {
	err := s.validateAddresses(to, from)
	if err != nil {
		return err
	}

	// implementation removed
	return nil
}

func (s Sender) validateAddresses(to, from string) error {
	// implementation removed
	return nil
}

type Message struct {
	Subject     string
	Message     string
	Attachments []os.File
}
