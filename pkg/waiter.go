package sim6502

import (
	"errors"
	"sync"
)

type waiter struct {
	lock    sync.Mutex
	waiting bool
	waiter  sync.WaitGroup
}

// wait, if not already waiting, will block until notify is called
// if already waiting, an error is returned
func (w *waiter) wait() error {
	w.lock.Lock()
	if w.waiting {
		w.lock.Unlock()
		return errors.New("already in wait")
	}
	w.waiting = true
	w.waiter.Add(1)
	w.lock.Unlock()
	w.waiter.Wait()
	return nil
}

// notify will:
// If there is a waiter, release it and return true
// if there is no waiter, return false
func (w *waiter) notify() bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.waiting {
		w.waiting = false
		w.waiter.Done()
		return true
	}
	return false
}

// isWaiting will return true if waiting
// really for testing only
func (w *waiter) isWaiting() bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.waiting
}
