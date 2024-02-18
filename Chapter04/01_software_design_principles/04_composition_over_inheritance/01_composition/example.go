package _1_composition

type bird struct{}

func (b bird) Fly() {
	// not implemented
}

type Duck struct {
	bird bird
}

func (d Duck) Fly() {
	d.bird.Fly()
}

func (d Duck) Talk() string {
	return "Quack!"
}

type Parrot struct {
	bird bird
}

func (p Parrot) Fly() {
	p.bird.Fly()
}

func (p Parrot) Talk() string {
	return "Squawk!"
}
