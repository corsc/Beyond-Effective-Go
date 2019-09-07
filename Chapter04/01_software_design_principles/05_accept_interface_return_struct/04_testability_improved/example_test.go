package _4_testability_improved

import (
	"testing"

	"github.com/PacktPublishing/Advanced-Go-Programming/Chapter04/01_software_design_principles/05_accept_interface_return_struct/04_testability_improved/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_happyPath(t *testing.T) {
	// mock the repository
	repository := &MockUserRepository{}
	repository.On("Save", mock.Anything).Return(nil)

	// build inputs
	testUser := &user.User{ID: 1, Name: "Amy", Email: "amy@home.com"}

	resultErr := repository.Save(testUser)
	assert.NoError(t, resultErr)
}
