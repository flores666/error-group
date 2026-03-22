package eg

import (
	"context"
	"sync"
)

type ErrorGroup struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	Error error
}

func NewErrorGroup() (*ErrorGroup, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())

	return &ErrorGroup{
		cancel: cancel,
		wg:     sync.WaitGroup{},
	}, ctx
}

func (e *ErrorGroup) Go(action func() error) {
	e.wg.Go(func() {
		if err := action(); err != nil {
			e.Error = err
			e.cancel()
			return
		}
	})
}

func (e *ErrorGroup) Wait() error {
	e.wg.Wait()
	return e.Error
}
