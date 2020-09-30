package _2_revisited

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	ErrBadRequest = errors.New("ID supplied is empty or the wrong format")
	ErrNotFound   = errors.New("user not found")
)

func loadUser(ID string) (*User, error) {
	userID, err := strconv.Atoi(ID)
	if err != nil {
		return nil, fmt.Errorf("%w - %s", ErrBadRequest, err)
	}

	user, err := loadFromDB(userID)
	if err != nil {
		return nil, fmt.Errorf("%w - userID: %d", ErrNotFound, userID)
	}

	return user, err
}

func Handler(resp http.ResponseWriter, req *http.Request) {
	userID := req.Form.Get("userID")

	user, err := loadUser(userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrBadRequest):
			resp.WriteHeader(http.StatusUnprocessableEntity)

		case errors.Is(err, ErrNotFound):
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

func loadFromDB(id int) (*User, error) {
	// not implemented
	return nil, errors.New("not implemented")
}
