package _0_maps

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// $ go test -v ./Chapter03/05_patterns/10_maps/ -run=Int
// === RUN   TestMapInt
// took: 22.555075ms
// took: 28.521478ms
// took: 24.487982ms
// took: 26.915654ms
// took: 26.770876ms
// --- PASS: TestMapInt (5.39s)
// PASS
// ok  	github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/10_maps	5.408s
// $ go test -v ./Chapter03/05_patterns/10_maps/ -run=String
// === RUN   TestMapString
// took: 32.853891ms
// took: 35.289389ms
// took: 34.021296ms
// took: 38.081171ms
// took: 35.818594ms
// --- PASS: TestMapString (5.73s)
// PASS
// ok  	github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/10_maps	5.750s
// $ go test -v ./Chapter03/05_patterns/10_maps/ -run=Person
// === RUN   TestMapPerson
// took: 22.676387ms
// took: 28.001048ms
// took: 26.874785ms
// took: 27.973287ms
// took: 27.39127ms
// --- PASS: TestMapPerson (5.40s)
// PASS
// ok  	github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/10_maps	5.424s
// $ go test -v ./Chapter03/05_patterns/10_maps/ -run=Pointer
// === RUN   TestMapPointer
// took: 49.99156ms
// took: 60.190347ms
// took: 51.53372ms
// took: 56.054801ms
// took: 53.660702ms
// --- PASS: TestMapPointer (5.70s)
// PASS
var dataInt = map[int]string{}

func TestMapInt(t *testing.T) {
	for x := 0; x < 1000000; x++ {
		dataInt[x] = "X"
	}

	for x := 0; x < 5; x++ {
		trackGC()
		time.Sleep(1 * time.Second)
	}
}

var dataString = map[string]string{}

func TestMapString(t *testing.T) {
	for x := 0; x < 1000000; x++ {
		dataString[strconv.Itoa(x)] = "X"
	}

	for x := 0; x < 5; x++ {
		trackGC()
		time.Sleep(1 * time.Second)
	}
}

var dataStruct = map[Person]string{}

func TestMapPerson(t *testing.T) {
	for x := 0; x < 1000000; x++ {
		dataStruct[Person{ID: x}] = "X"
	}

	for x := 0; x < 5; x++ {
		trackGC()
		time.Sleep(1 * time.Second)
	}
}

var dataPointer = map[*Person]string{}

func TestMapPointer(t *testing.T) {
	for x := 0; x < 1000000; x++ {
		dataPointer[&Person{ID: x}] = "X"
	}

	for x := 0; x < 5; x++ {
		trackGC()
		time.Sleep(1 * time.Second)
	}
}

type Person struct {
	ID int
}

func trackGC() {
	t := time.Now()
	runtime.GC()
	fmt.Printf("took: %s\n", time.Since(t))
}
