package _2_named_errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrBadRequest    = errors.New("ID supplied is empty or the wrong format")
	ErrNotFound      = errors.New("user not found")
	ErrNotAuthorized = errors.New("permission denied")
)

func loadUser(ID string) (*User, error) {
	// implementation removed
	return nil, errors.New("not implemented")
}

func Handler(resp http.ResponseWriter, req *http.Request) {
	userID := req.Form.Get("userID")

	user, err := loadUser(userID)
	if err != nil {
		switch err {
		case ErrBadRequest:
			resp.WriteHeader(http.StatusUnprocessableEntity)

		case ErrNotAuthorized:
			resp.WriteHeader(http.StatusUnauthorized)

		case ErrNotFound:
			resp.WriteHeader(http.StatusNotFound)

		default:
			resp.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	encoder := json.NewEncoder(resp)
	_ = encoder.Encode(user)
}

type User struct {
	// implementation removed
}
