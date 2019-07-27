package _5_extended

import (
	"bytes"
	"strconv"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func toString(people []Person) string {
	buffer := bufferPool.Get().(*bytes.Buffer)

	for _, person := range people {
		buffer.WriteString("ID: ")
		buffer.WriteString(strconv.Itoa(person.ID))
		buffer.WriteString("\nName: ")
		buffer.WriteString(person.Name)
	}

	result := buffer.String()

	buffer.Reset()
	bufferPool.Put(buffer)

	return result
}

type Person struct {
	ID   int
	Name string
}
