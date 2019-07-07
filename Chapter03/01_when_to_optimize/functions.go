package _1_when_to_optimize

import (
	"bytes"
	"fmt"
	"strconv"
)

func CleanExample(person Person) string {
	return fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
}

func FastExample(buffer *bytes.Buffer, person Person) string {
	buffer.Reset()

	_, _ = buffer.WriteString("ID: ")
	_, _ = buffer.WriteString(strconv.FormatInt(person.ID, 10))
	_, _ = buffer.WriteString("\nName: ")
	_, _ = buffer.WriteString(person.Name)

	return buffer.String()
}

type Person struct {
	ID   int64
	Name string
}
