package _1_abstract_methods

func ExampleImage() {
	myPNG := &Image{
		toBytes: pngEncoder,
	}

	myPNG.Save("my-file.png")
}

func pngEncoder() []byte {
	// implementation removed
	return nil
}
