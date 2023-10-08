package lock

import (
	"context"
	"errors"
	"time"
)

type Lock interface {
	Lock()
	TryLock() bool
	TryLockWithTimeout(timeout time.Duration) bool
	Unlock()
}

type lock struct {
	lock chan struct{}
}

// Lock put an element into channel to mark that locked
func (l *lock) Lock() {
	l.lock <- struct{}{}
}

// TryLock returns true if trying lock success else false
func (l *lock) TryLock() (locked bool) {
	select {
	case l.lock <- struct{}{}:
		locked = true
	default:
	}
	return
}

// TryLockWithTimeout returns true if trying lock success that within the specified time else false
func (l *lock) TryLockWithTimeout(timeout time.Duration) (locked bool) {
	timeoutCtx, cancelFunc := context.WithTimeout(context.TODO(), timeout)
	defer cancelFunc()
	select {
	case l.lock <- struct{}{}:
		locked = true
	case <-timeoutCtx.Done():
	}
	return
}

// Unlock panic if unlock of unlocked lock
func (l *lock) Unlock() {
	select {
	case <-l.lock:
	default:
		panic(errors.New("unlock of unlocked lock"))
	}
	return
}

func New() Lock {
	return &lock{
		lock: make(chan struct{}, 1),
	}
}
