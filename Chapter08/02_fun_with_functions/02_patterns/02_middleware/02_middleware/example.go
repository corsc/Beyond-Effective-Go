package _2_middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Usage() {
	listHandler := http.HandlerFunc(ListUsersHandler)
	http.Handle("/", trackRequest(listHandler))
	http.ListenAndServe(":8080", nil)
}

func ListUsersHandler(resp http.ResponseWriter, req *http.Request) {
	users := loadAllUsers()

	payload, _ := json.Marshal(users)
	_, _ = resp.Write(payload)
}

func trackRequest(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()

		handler(resp, req)

		log.Printf("Time taken: %s", time.Since(start))
	})
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
