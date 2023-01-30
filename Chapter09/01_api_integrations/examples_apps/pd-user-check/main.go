package main

import (
	"flag"
	"fmt"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pd-user-check/internal/pduty"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pd-user-check/internal/printer"
	"log"
	"os"
)

var (
	onlyErrors bool

	requireEmail bool
	requireSMS   bool
	requirePhone bool
	requirePush  bool
)

func main() {
	// these are inputs that should come from command line or env
	apiKey, found := os.LookupEnv("PD_API_KEY")
	if !found {
		log.Fatal("PD_API_KEY must be set")
	}

	flag.BoolVar(&onlyErrors, "errors", true, "print only those users with invalid settings")
	flag.BoolVar(&requireEmail, "email", true, "require Email setting")
	flag.BoolVar(&requireSMS, "sms", false, "require SMS setting")
	flag.BoolVar(&requirePhone, "phone", true, "require Phone setting")
	flag.BoolVar(&requirePush, "push", false, "require Push setting")
	flag.Parse()

	// actual logic
	fmt.Printf("Fetching all pduty\n")
	users, err := (&pduty.UsersAPI{}).GetUsers(apiKey, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Validating\n")
	validator := &printer.PrinterAPI{
		RequireEmail: requireEmail,
		RequireSMS:   requireSMS,
		RequirePhone: requirePhone,
		RequirePush:  requirePush,
	}
	validator.OutputHeader(os.Stdout)

	for _, user := range users {
		err := validator.Print(os.Stdout,
			&printer.User{
				ID:       user.ID,
				Name:     user.Name,
				Email:    user.Email,
				Teams:    user.Teams,
				EmailSet: user.EmailSet,
				PhoneSet: user.PhoneSet,
				SMSSet:   user.SMSSet,
				PushSet:  user.PushSet,
			}, onlyErrors)

		if err != nil {
			log.Fatal(err)
		}
	}
}
