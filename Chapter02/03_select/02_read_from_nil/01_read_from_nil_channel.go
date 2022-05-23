package _2_read_from_nil

func readFromNilChannel() []string {
	var dataCh chan struct{}
	var results []string

	for x := 0; x < 4; x++ {
		select {
		case <-dataCh:
			results = append(results, "read from channel")

			// nil the channel
			dataCh = nil

		default:
			results = append(results, "default")

			// restore the channel
			dataCh = make(chan struct{}, 1)
			dataCh <- struct{}{}
		}
	}

	// Output here is:
	// []string{
	//	"default",
	//	"read from channel",
	//	"default",
	//	"read from channel",
	// }
	return results
}
