package _1_accept_interface

func SendEmail(recipient Recipient, subject, message string) error {
	// implementation removed
	return nil
}

type Recipient interface {
	GetEmailAddress() string
}
