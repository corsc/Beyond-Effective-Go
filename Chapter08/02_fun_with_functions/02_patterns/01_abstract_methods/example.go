package _1_abstract_methods

import (
	"os"
)

type Image struct {
	toBytes func() []byte
}

func (i *Image) Save(filename string) {
	destination := i.openFile(filename)
	data := i.toBytes()
	i.writeToFile(data, destination)
	i.closeFile(destination)
}

func (i *Image) openFile(filename string) *os.File {
	// implementation removed
	return nil
}

func (i *Image) writeToFile(data []byte, file *os.File) {
	// implementation removed
}

func (i *Image) closeFile(file *os.File) {
	// implementation removed
}
