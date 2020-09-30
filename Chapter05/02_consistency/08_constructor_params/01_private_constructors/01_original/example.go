package _1_original

import (
	"time"
)

func NewUserManager(minPwdLen, maxPwdLen int, pwdEncoder PasswordEncoder, storage Storage, now func() time.Time, risk RiskManager) *UserManager {
	return &UserManager{
		minPwdLen:  minPwdLen,
		maxPwdLen:  maxPwdLen,
		pwdEncoder: pwdEncoder,
		storage:    storage,
		now:        now,
		risk:       risk,
	}
}

type UserManager struct {
	minPwdLen  int
	maxPwdLen  int
	pwdEncoder PasswordEncoder
	storage    Storage
	now        func() time.Time
	risk       RiskManager
}

type PasswordEncoder interface {
	Encode(plainText string) (cipherText string)
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
