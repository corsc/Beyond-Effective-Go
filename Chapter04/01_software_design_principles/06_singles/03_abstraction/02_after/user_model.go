package _2_after

type UserModel struct {
	dao *UserDAO
}

// Attempts to create the supplied user
func (u *UserModel) Create(user *User) (int, error) {
	err := u.validate(user)
	if err != nil {
		return 0, err
	}

	return u.dao.Save(user)
}

// validate the supplied user is complete and correct
func (u *UserModel) validate(user *User) error {
	// implementation removed
	return nil
}
