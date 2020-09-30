package _3_whitespace

func doCreateUser(req *createUser) error {
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

func validate(req *createUser) error {
	// implementation removed
	return nil
}

func saveToDB(req *createUser) error {
	// implementation removed
	return nil
}
