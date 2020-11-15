package _3_quick_mocks

type Crud interface {
	Save(u *User) error
	Update(u *User) error
	Delete(u *User) error
}

type MockSave struct {
	Crud

	OnSave func(u *User) error
}

func (m *MockSave) Save(u *User) error {
	return m.OnSave(u)
}

type User struct {
	// fields removed
}
