package _1_accept_interface

func SendEmail(recipient Recipient, title, message string) error {
	// implementation removed
	return nil
}

type Recipient interface {
	EmailAddress() string
}
