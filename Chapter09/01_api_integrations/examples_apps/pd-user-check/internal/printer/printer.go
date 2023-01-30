package printer

import (
	"fmt"
	"io"
)

// PrinterAPI validates the supplied user against the required settings and (optionally sends notifications)
type PrinterAPI struct {
	RequireEmail bool
	RequireSMS   bool
	RequirePhone bool
	RequirePush  bool
}

// Print will output the supplied user either always or only when in error (depends on settings)
func (v *PrinterAPI) Print(logger io.Writer, user *User, showOnlyErrors bool) error {
	if showOnlyErrors {
		if !v.isValid(user) {
			v.outputLine(logger, user)
		}
	} else {
		// always output
		v.outputLine(logger, user)
	}

	return nil
}

func (v *PrinterAPI) outputLine(logger io.Writer, user *User) {
	_, _ = fmt.Fprintf(logger, "%-40s %-60s %-8d %-8t %-8t %-8t %-8t\n", user.Name, user.Email, user.Teams, user.EmailSet, user.PhoneSet, user.SMSSet, user.PushSet)
}

// OutputHeader outputs the header of the result table
func (v *PrinterAPI) OutputHeader(logger io.Writer) {
	_, _ = fmt.Fprintf(logger, "%-40s %-60s %-8s %-8s %-8s %-8s %-8s\n", "Name", "Email", "Teams", "Email", "Phone", "SMS", "Push")
	_, _ = fmt.Fprintf(logger, "-----------------------------------------------------------------------------------------------------------------------------------------------------\n")
}

func (v *PrinterAPI) isValid(user *User) bool {
	if v.RequireEmail && !user.EmailSet {
		return false
	}
	if v.RequireSMS && !user.SMSSet {
		return false
	}
	if v.RequirePhone && !user.PhoneSet {
		return false
	}
	if v.RequirePush && !user.PushSet {
		return false
	}

	return true
}

type User struct {
	ID       string
	Name     string
	Email    string
	Teams    int
	EmailSet bool
	PhoneSet bool
	SMSSet   bool
	PushSet  bool
}
