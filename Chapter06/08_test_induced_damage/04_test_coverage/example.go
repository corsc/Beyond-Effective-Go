package _4_test_coverage

import (
	"encoding/json"
	"fmt"
)

func UserAsJSON(name, address string, age int) ([]byte, error) {
	user := &User{
		Name:    name,
		Address: address,
		Age:     age,
	}

	payload, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to convert User to JSON with err: %w", err)
	}

	return payload, nil
}

type User struct {
	Name    string
	Address string
	Age     int
}
