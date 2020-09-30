package _1_exceptions

import (
	"fmt"
	"strconv"
)

func LoadUser(idAsString string) (out *User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to load user with err :%s", r)
		}
	}()

	id := parseID(idAsString)
	out = loadByID(id)

	return
}

func parseID(idAsString string) int {
	out, err := strconv.Atoi(idAsString)
	if err != nil {
		panic(err)
	}

	return out
}

func loadByID(id int) *User {
	// implementation removed
	return nil
}

type User struct {
	// implementation removed
}
