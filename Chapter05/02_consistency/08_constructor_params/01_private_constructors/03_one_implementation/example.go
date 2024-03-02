package example

import (
	"time"
)

func NewUserManager(minNameLen, maxNameLen int, storage Storage, risk RiskManager) *UserManager {
	return newUserManager(minNameLen, maxNameLen, &myPasswordEncoder{}, storage, time.Now, risk)
}

func newUserManager(minNameLen, maxNameLen int, pwdEncoder PasswordEncoder, storage Storage, now func() time.Time, risk RiskManager) *UserManager {
	return &UserManager{
		minNameLen: minNameLen,
		maxNameLen: maxNameLen,
		pwdEncoder: pwdEncoder,
		storage:    storage,
		now:        now,
		risk:       risk,
	}
}

type UserManager struct {
	minNameLen int
	maxNameLen int
	pwdEncoder PasswordEncoder
	storage    Storage
	now        func() time.Time
	risk       interface{}
}

type PasswordEncoder interface {
	Encode(plainText string) (cipherText string)
}

// lone implementation of the password encoder
type myPasswordEncoder struct{}

func (m *myPasswordEncoder) Encode(plainText string) (cipherText string) {
	// implementation removed
	return ""
}

type Storage interface {
	LoadByID(ID int) (User, error)
	Save(user User) error
}

type User struct {
	// implementation removed
}

type RiskManager interface {
	Check(user User) error
}
