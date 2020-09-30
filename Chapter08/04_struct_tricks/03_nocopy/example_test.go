package _3_nocopy

type Cache struct {
	// implementation removed

	noCopy
}

type noCopy struct{}

func (*noCopy) Lock() {}

func (*noCopy) Unlock() {}

func Example() {
	myCache := Cache{}

	// attempt to pass by value
	useCache(myCache)
}

func useCache(cache Cache) {
	// implementation removed
}
