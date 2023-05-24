package gonce

import (
	"sync"
	"sync/atomic"
)

type Once[T any] struct {
	result T
	mu     sync.Mutex
	done   uint32
}

// Do function executes a given function:
// - if the function returned an error, then it can be called again;
// - if no error returned, then the returned result will be stored and used later;
// - if the function has already been executed, then the stored result will be returned.
func (o *Once[T]) Do(f func() (result T, err error)) (result T, err error) {
	if atomic.LoadUint32(&o.done) == 1 {
		return o.result, nil
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		result, err = f()
		if err != nil {
			return
		}
		o.result = result
		atomic.StoreUint32(&o.done, 1)
	}
	return o.result, nil
}
