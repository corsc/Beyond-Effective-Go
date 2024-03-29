package _3_output

import (
	"fmt"
	"io"
)

type Printer struct {
	logger io.Writer
}

func (p *Printer) outputHeader() {
	_, _ = fmt.Fprintf(p.logger, "%-40s %-60s %-8s %-8s %-8s %-8s\n",
		"Name", "Email Address", "Email", "Phone", "SMS", "Push")
	_, _ = fmt.Fprint(p.logger,
		"-----------------------------------------------------------------------------------------------------------------------------------------------------\n")
}

func (p *Printer) outputLine(user *User) {
	_, _ = fmt.Fprintf(p.logger, "%-40s %-60s %-8t %-8t %-8t %-8t\n",
		user.Name, user.Email, user.EmailIsSet, user.PhoneIsSet,
		user.SMSIsSet, user.PushIsSet)
}

type User struct {
	ID         string
	Name       string
	Email      string
	EmailIsSet bool
	PhoneIsSet bool
	SMSIsSet   bool
	PushIsSet  bool
}
