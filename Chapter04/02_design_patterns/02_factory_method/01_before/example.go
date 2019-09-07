package _1_before

func SingSolo(favorite string) {
	var beatle Beatle

	switch favorite {
	case "John":
		beatle = John{}

	case "Paul":
		beatle = Paul{}

	case "George":
		beatle = George{}

	case "Ringo":
		beatle = Ringo{}
	}

	beatle.Sing()
}

type Beatle interface {
	Sing()
	PlayInstrument()
}

type John struct{}

func (j John) Sing() {
	// implementation removed
}

func (j John) PlayInstrument() {
	// implementation removed
}

type Paul struct{}

func (j Paul) Sing() {
	// implementation removed
}

func (j Paul) PlayInstrument() {
	// implementation removed
}

type George struct{}

func (j George) Sing() {
	// implementation removed
}

func (j George) PlayInstrument() {
	// implementation removed
}

type Ringo struct{}

func (j Ringo) Sing() {
	// implementation removed
}

func (j Ringo) PlayInstrument() {
	// implementation removed
}
