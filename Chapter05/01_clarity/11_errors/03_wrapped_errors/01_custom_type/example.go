package _1_custom_type

type BadRequestError struct {
	field string
}

func (b *BadRequestError) Error() string {
	return b.field + " was missing or invalid"
}

func Usage(err error) {
	if err != nil {
		if _, ok := err.(*BadRequestError); ok {
			// bad request
			return
		}

		// other errors
	}
}
