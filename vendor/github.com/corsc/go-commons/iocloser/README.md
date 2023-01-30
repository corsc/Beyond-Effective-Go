# iocloser 

This package intends to provide an easy and concise of closing io.Closers.

Typically you will find close that looks like:

    defer reader.Close()

Your lint checkers will likely start complaining that you didn't properly handle the error.  An error that you really don't care about.

So you might decide to do:

    defer func() {
 		_ = reader.Close()
	  }()

This will make some linters happy but others will not let you get off so easily.

So instead of writing the following (arguably even uglier code):

    defer func() {
 		err := reader.Close()
		if err != nil {
			log.Printf("error was: %v", err)
		}
	  }()

This package allows you to get it back to one line:

    defer iocloser.Close(reader, log.Printf)

Or just

    defer iocloser.Close(reader)

For usage examples please refer [here](iocloser_examples_test.go)