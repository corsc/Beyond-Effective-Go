package _1_no_middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Usage() {
	listHandler := http.HandlerFunc(ListUsersHandler)
	http.Handle("/", listHandler)
	http.ListenAndServe(":8080", nil)
}

func ListUsersHandler(resp http.ResponseWriter, req *http.Request) {
	start := time.Now()

	users := loadAllUsers()

	payload, _ := json.Marshal(users)
	_, _ = resp.Write(payload)

	log.Printf("Time taken: %s", time.Since(start))
}

func loadAllUsers() *Users {
	return nil
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
}
