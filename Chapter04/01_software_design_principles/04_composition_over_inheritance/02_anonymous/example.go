package _2_anonymous

type bird struct{}

func (b bird) Fly() {
	// not implemented
}

type Duck struct {
	bird
}

func (d *Duck) Talk() string {
	return "Quack!"
}

type Parrot struct {
	bird
}

func (p Parrot) Talk() string {
	return "Squawk!"
}
