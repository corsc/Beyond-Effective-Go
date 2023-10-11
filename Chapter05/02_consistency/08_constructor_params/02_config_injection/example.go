package example

import (
	"database/sql"
	"time"
)

func NewUserManager(cfg Config, storage Storage, risk RiskManager) *UserManager {
	return newUserManager(cfg, &myPasswordEncoder{}, storage, time.Now, risk)
}

func newUserManager(cfg Config, pwdEncoder PasswordEncoder, storage Storage, now func() time.Time, risk RiskManager) *UserManager {
	return &UserManager{
		cfg:        cfg,
		pwdEncoder: pwdEncoder,
		storage:    storage,
		now:        now,
		risk:       risk,
	}
}

type UserManager struct {
	cfg        Config
	pwdEncoder PasswordEncoder
	storage    Storage
	now        func() time.Time
	risk       interface{}
}

type Config interface {
	GetDBPool() *sql.DB

	GetMinPwdLen() int
	GetMaxPwdLen() int
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
