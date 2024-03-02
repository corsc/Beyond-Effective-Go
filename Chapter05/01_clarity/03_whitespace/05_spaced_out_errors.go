package _3_whitespace

func doCreateUser(req *userCreationRequest) error {
	err := validate(req)

	if err != nil {
		return err
	}

	err = saveToDB(req)

	if err != nil {
		return err
	}

	return nil
}

func validate(req *userCreationRequest) error {
	// implementation removed
	return nil
}

func saveToDB(req *userCreationRequest) error {
	// implementation removed
	return nil
}
