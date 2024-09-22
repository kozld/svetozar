package workerpool

// Step 1: Define the Task
// A task that accepts a URL and returns the extracted data as a string.
type Task interface {
	Run()
}

// Step 2: Create the Worker
// A worker is a goroutine that processes tasks and sends the results through a channel.
type Worker struct {
	id        int
	taskQueue <-chan Task
}

func (w *Worker) Start() {
	go func() {
		for task := range w.taskQueue {
			fetchAndProcess(task)
		}
	}()
}

// Step 3: Implement the Worker Pool
// The worker pool manages the workers, distributes tasks, and collects results.
type WorkerPool struct {
	taskQueue   chan Task
	workerCount int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan Task),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{id: i, taskQueue: wp.taskQueue}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.taskQueue <- task
}

// Fetch and process the task
func fetchAndProcess(task Task) {
	task.Run()
}
