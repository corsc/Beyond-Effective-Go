package _3_signal

func NewWorkerPool(total int) *WorkerPool {
	workerPool := &WorkerPool{
		shutdownCh: make(chan struct{}),
	}

	for x := 0; x < total; x++ {
		go doWork(workerPool.shutdownCh)
	}

	return workerPool
}

type WorkerPool struct {
	shutdownCh chan struct{}
}

func doWork(shutdownCh chan struct{}) {
	for {
		select {
		case <-shutdownCh:
			// quit
			return

		default:
		}

		// do work - implementation remove
	}
}

func (w *WorkerPool) Shutdown() {
	close(w.shutdownCh)
}
