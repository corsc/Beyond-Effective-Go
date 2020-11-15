package _1_untestable

import (
	"encoding/json"
	"net/http"
)

func GetUserAPI(resp http.ResponseWriter, req *http.Request) {
	ID := getID(req)

	user := loadUser(ID)

	payload, err := json.Marshal(user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = resp.Write(payload)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getID(req *http.Request) int64 {
	return 666
}

func loadUser(id int64) *User {
	return &User{
		ID:   666,
		Name: "Sophia",
	}
}

type User struct {
	ID   int64
	Name string
}
