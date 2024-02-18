package _2_after

import (
	"net/http"
	"strconv"
)

type CreateUserHandler struct {
	model *UserModel
}

func (s *CreateUserHandler) Handler(resp http.ResponseWriter, req *http.Request) {
	user, err := s.extractUser(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.model.Create(user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("location", "/user/"+strconv.Itoa(id))
}

// extract user from HTTP request
func (s *CreateUserHandler) extractUser(request *http.Request) (*User, error) {
	// implementation removed
	return &User{}, nil
}

type User struct {
	Name string
}
