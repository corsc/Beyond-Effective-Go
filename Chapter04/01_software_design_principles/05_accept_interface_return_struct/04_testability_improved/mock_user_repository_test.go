// Code generated by mockery v1.0.0. DO NOT EDIT.

package _4_testability_improved

import (
	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/04_testability_improved/internal/user"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: _a0
func (_m *MockUserRepository) Save(_a0 *user.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
