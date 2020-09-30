package _7_shadowing

func performTask() error {
	resultCh := make(chan error)

	err := doWork(resultCh)
	if err != nil {
		return err
	}

	for err := range resultCh {
		handleError(err)
	}

	return err
}

func doWork(resultCh chan error) error {
	// implementation removed
	return nil
}

func handleError(err error) {
	// implementation removed
}
