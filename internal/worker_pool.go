package internal

type WorkerPool struct {
	workersCount int
	pool         chan *Worker
	results      chan HideURIsResult
}

func NewWorkerPool(workersCount int, handler func(string) string) *WorkerPool {
	pool := make(chan *Worker, workersCount)
	for i := range workersCount {
		pool <- NewWorker(i, handler)
	}

	return &WorkerPool{
		workersCount: workersCount,
		pool:         pool,
		results:      make(chan HideURIsResult),
	}
}

func (p *WorkerPool) Process(jobs <-chan HideURIsJob) {
	for job := range jobs {
		w := <-p.pool
		go func(job HideURIsJob, w *Worker) {
			p.results <- HideURIsResult{job.index, w.handler(job.str)}
			p.pool <- w
		}(job, w)
	}
	p.wait()
}

func (p *WorkerPool) Results() <-chan HideURIsResult {
	return p.results
}

func (p *WorkerPool) wait() {
	for range p.workersCount {
		<-p.pool
	}
	close(p.results)
}
