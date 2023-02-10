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
	Unlock() error
}

type lock struct {
	lock chan struct{}
}

func (l *lock) Lock() {
	l.lock <- struct{}{}
}

func (l *lock) TryLock() (locked bool) {
	select {
	case l.lock <- struct{}{}:
		locked = true
	default:
	}
	return
}

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

func (l *lock) Unlock() (err error) {
	select {
	case <-l.lock:
	default:
		err = errors.New("unlock of unlocked lock")
	}
	return
}

func New() Lock {
	return &lock{
		lock: make(chan struct{}, 1),
	}
}
