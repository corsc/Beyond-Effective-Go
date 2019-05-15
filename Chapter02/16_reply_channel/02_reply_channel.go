package _6_reply_channel

func CloseReplyChannel() []error {
	errorCh := make(chan error, 3)

	// call close on everything
	go CloseA(errorCh)
	go CloseB(errorCh)
	go CloseC(errorCh)

	// collate the errors
	var errors []error
	for x := 0; x < 3; x++ {
		err := <-errorCh
		errors = append(errors, err)
	}

	return errors
}

func CloseA(errorCh chan error) {
	err := doClose()

	// send result
	errorCh <- err
}

func CloseB(errorCh chan error) {
	// perform close
	var err error

	// send result
	errorCh <- err
}

func CloseC(errorCh chan error) {
	// perform close
	var err error

	// send result
	errorCh <- err
}
