package _3_whitespace

func doCreateUser(req *createUserRequest) error {
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

func validate(req *createUserRequest) error {
	// implementation removed
	return nil
}

func saveToDB(req *createUserRequest) error {
	// implementation removed
	return nil
}
