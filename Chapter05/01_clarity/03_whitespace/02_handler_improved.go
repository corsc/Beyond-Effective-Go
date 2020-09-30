package _3_whitespace

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handlerWhitespace(resp http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
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
