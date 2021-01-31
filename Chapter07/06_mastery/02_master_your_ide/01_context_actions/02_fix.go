package _1_context_actions

import (
	"encoding/json"
)

func Example2(payload []byte) {

	user := &User{}

	json.Unmarshal(payload, user)

}
