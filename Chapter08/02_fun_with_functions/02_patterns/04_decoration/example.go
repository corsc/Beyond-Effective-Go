package _4_decoration

import (
	"net/http"
)

func Usage() {
	var myHandler http.Handler = http.HandlerFunc(LoginHandler)

	http.Handle("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

func LoginHandler(resp http.ResponseWriter, req *http.Request) {
	// implementation removed
}
