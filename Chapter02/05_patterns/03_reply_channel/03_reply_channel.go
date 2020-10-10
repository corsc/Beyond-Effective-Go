package _3_reply_channel

type Workers struct {
	workCh chan Work
}

func worker(workCh chan Work) {
	for work := range workCh {
		result := sum(work)

		work.replyCh <- result
	}
}

func (w *Workers) SubmitWork(work Work) {
	w.workCh <- work
}

func (w *Workers) startWorker(workers int) {
	for x := 0; x < workers; x++ {
		go worker(w.workCh)
	}
}

func sum(work Work) int {
	return work.inputA + work.inputB
}

type Work struct {
	inputA int
	inputB int

	replyCh chan int
}
