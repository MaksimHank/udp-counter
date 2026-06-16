package workerPool

import (
	"context"
	"sync"
)

type Pool struct {
	tasks  chan func() error
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}
