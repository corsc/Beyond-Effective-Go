package _3_builder

import (
	"errors"
)

type FileBuilder struct{}

func (f *FileBuilder) Build(version int, encoder Encoder, decoder Decoder) (File, error) {
	switch version {
	case 1:
		return &fileV1{encoder: encoder, decoder: decoder}, nil

	default:
		return nil, errors.New("unknown version")
	}
}

type File interface {
	Save(filename string) error
}

type fileV1 struct {
	encoder Encoder
	decoder Decoder
}

func (f *fileV1) Save(filename string) error {
	// implementation removed
	return nil
}

type Encoder interface {
	// implementation removed
}

type Decoder interface {
	// implementation removed
}
