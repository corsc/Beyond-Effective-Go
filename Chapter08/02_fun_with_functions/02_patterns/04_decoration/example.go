package _4_decoration

import (
	"net/http"
)

func StartServer() {
	var myHandler http.Handler = MyHandlerFunc(SayHelloHandler)

	http.Handle("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

type MyHandlerFunc func(resp http.ResponseWriter, req *http.Request)

func (m MyHandlerFunc) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	m(resp, req)
}

func SayHelloHandler(resp http.ResponseWriter, req *http.Request) {
	_, _ = resp.Write([]byte(`Hello World!`))
}
