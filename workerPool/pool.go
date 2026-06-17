package workerPool

import (
	"context"
	"sync"
)

type Pool struct {
	tasks      chan func() error
	NumWorkers int
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

func NewPool(numWorkers int) *Pool {
	tasks := make(chan func() error, numWorkers)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	return &Pool{
		tasks:      tasks,
		NumWorkers: numWorkers,
		cancel:     cancel,
		ctx:        ctx,
		wg:         wg,
	}
}

func (p *Pool) Run(n int) {
	for n = range p.NumWorkers {
		wg.Add(1)
		go worker(n, p.tasks, p.ctx, &p.wg)
	}
}
