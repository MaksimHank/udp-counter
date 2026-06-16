package workerPool

func worker(id int, tasks <-chan func() error, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-tasks:
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
