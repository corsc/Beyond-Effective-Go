package _1_style

func ErrorHandling() error {
	// compressed style
	if err := doTask(); err != nil {
		return err
	}

	// longer style
	err := doTask()
	if err != nil {
		return err
	}

	return nil
}

func doTask() error {
	// implementation removed
	return nil
}
