package workerPool

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type Pool struct {
	tasks      chan func() error
	numWorkers int
	ctx        context.Context
	cancel     context.CancelFunc
	wg         *sync.WaitGroup
}

func NewPool(numWorkers int) *Pool {
	tasks := make(chan func() error, numWorkers)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	return &Pool{
		tasks:      tasks,
		numWorkers: numWorkers,
		cancel:     cancel,
		ctx:        ctx,
		wg:         wg,
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go worker(i, p.tasks, p.ctx, p.wg)
	}
}

func (p *Pool) Submit(task func() error) error {
	select {
	case <-p.ctx.Done():
		fmt.Printf("func Submit. \nPool is shutting down, task rejected")
		return errors.New("func Submit. \nPool is shutting down, task rejected")
	case p.tasks <- task:
		fmt.Printf("func Submit. \nTask was successfully sent")
		return nil
	}
}

func (p *Pool) Shutdown() {
	p.cancel()
	close(p.tasks)
	p.wg.Wait()
}
