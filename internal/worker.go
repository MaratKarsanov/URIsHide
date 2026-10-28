package internal

type Worker struct {
	id      int
	handler func(string) string
}

func NewWorker(id int, handler func(string) string) *Worker {
	return &Worker{
		id:      id,
		handler: handler,
	}
}
