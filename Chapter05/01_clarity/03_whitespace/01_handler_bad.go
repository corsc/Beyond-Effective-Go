package _3_whitespace

import (
	"encoding/json"
	"io"
	"net/http"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	createReq := &createUser{}
	err = json.Unmarshal(payload, createReq)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	err = doCreateUser(createReq)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = resp.Write([]byte("OK"))
}

type createUser struct {
	// fields removed
}
