package _3_reply_channel

type WorkerGroup struct {
	taskCh chan Task
}

func NewWorkers(totalWorkers int) *WorkerGroup {
	workerGroup := &WorkerGroup{
		taskCh: make(chan Task, 10),
	}

	for x := 0; x < totalWorkers; x++ {
		go workerGroup.newWorker()
	}

	return workerGroup
}

func (w *WorkerGroup) newWorker() {
	for newTask := range w.taskCh {
		// this will typically be a heavy operation
		newTask.Do()
	}
}

func (w *WorkerGroup) SubmitTask(in Task) {
	w.taskCh <- in
}

type Task struct {
	inputA int
	inputB int

	replyCh chan int
}

func (w *Task) Do() {
	w.replyCh <- w.inputA + w.inputB
}
