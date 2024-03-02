package _1_meaningful

func Example_structNames() {
	// Bad
	type cur struct {
		// fields removed
	}

	// Better
	type create struct {
		// fields removed
	}

	// Best
	type userCreationRequest struct {
		// fields removed
	}
}
