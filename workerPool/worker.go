package workerPool

import (
	"context"
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan func() error, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				wg.Done()
				return
			}
			err := task()
			if err != nil {
				fmt.Printf("%v: error from worker %d", err, id)
			}
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
