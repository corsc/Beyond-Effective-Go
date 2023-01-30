package main

import (
	"flag"
)

func main() {
	var (
		onlyErrors bool

		requireEmail bool
		requireSMS   bool
		requirePhone bool
		requirePush  bool
	)

	flag.BoolVar(&onlyErrors, "errors", true, "print only those users with invalid settings")
	flag.BoolVar(&requireEmail, "email", true, "require Email setting")
	flag.BoolVar(&requireSMS, "sms", false, "require SMS setting")
	flag.BoolVar(&requirePhone, "phone", true, "require Phone setting")
	flag.BoolVar(&requirePush, "push", false, "require Push setting")
	flag.Parse()

	// call the business logic here
}
