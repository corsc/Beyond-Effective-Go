package _6_test_resilience

type Unit struct {
	a DepA
	b DepB
	c DepC
}

func (u *Unit) Do() error {
	err := u.a.Do()
	if err != nil {
		return err
	}

	err = u.b.Do()
	if err != nil {
		return err
	}

	return u.c.Do()
}

//go:generate mockery --name=DepA --case underscore --testonly --inpackage
type DepA interface {
	Do() error
}

//go:generate mockery --name=DepB --case underscore --testonly --inpackage
type DepB interface {
	Do() error
}

//go:generate mockery --name=DepC --case underscore --testonly --inpackage
type DepC interface {
	Do() error
}
