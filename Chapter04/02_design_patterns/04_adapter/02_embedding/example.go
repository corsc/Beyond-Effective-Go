package _2_embedding

type User struct {
	username string
	password string
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

type UserForLogin interface {
	Username() string
	Password() string
	Token() string
}

type UserLoginAdapter struct {
	User
}

func (u *UserLoginAdapter) Token() string {
	// implementation removed
	return ""
}

// Enforce the relationship between the adapter and the new interface.
var _ UserForLogin = &UserLoginAdapter{User: User{}}
