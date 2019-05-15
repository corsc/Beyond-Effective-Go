package _6_reply_channel

func CloseTraditional() []error {
	// call close on everything
	errCh1 := Close1()
	errCh2 := Close2()
	errCh3 := Close3()

	// collate responses
	var errors []error

	err := <-errCh1
	if err != nil {
		errors = append(errors, err)
	}

	err = <-errCh2
	if err != nil {
		errors = append(errors, err)
	}

	err = <-errCh3
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}

func Close1() chan error {
	errorCh := make(chan error, 1)

	go func() {
		err := doClose()

		// send result
		errorCh <- err
	}()

	return errorCh
}

func doClose() error {
	return nil
}

func Close2() chan error {
	errorCh := make(chan error, 1)

	go func() {
		// perform close
		var err error

		// send result
		errorCh <- err
	}()

	return errorCh
}

func Close3() chan error {
	errorCh := make(chan error, 1)

	go func() {
		// perform close
		var err error

		// send result
		errorCh <- err
	}()

	return errorCh
}
