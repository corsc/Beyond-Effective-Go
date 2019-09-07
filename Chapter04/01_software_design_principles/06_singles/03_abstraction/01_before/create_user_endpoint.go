package _1_before

import (
	"net/http"
	"strconv"
)

type CreateUserEndpoint struct {
	validator *Validator
	dao       *UserDAO
}

func (s *CreateUserEndpoint) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	user, err := s.extractUser(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.validator.Validate(user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.dao.Save(user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("location", "/user/"+strconv.Itoa(id))
}

// extract user from HTTP request
func (s *CreateUserEndpoint) extractUser(request *http.Request) (*User, error) {
	// implementation removed
	return &User{}, nil
}

type User struct {
	Name string
}
