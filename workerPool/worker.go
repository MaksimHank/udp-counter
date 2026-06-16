package workerPool

func worker(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n
	}
}
