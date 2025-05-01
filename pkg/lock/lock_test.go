package lock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLock_TryLock(t *testing.T) {
	ast := assert.New(t)
	l := New()

	locked := l.TryLock()
	ast.True(locked)
	l.Unlock()

	l.Lock()
	locked = l.TryLock()
	ast.False(locked)
	l.Unlock()
}

func TestLock_TryLockWithTimeout(t *testing.T) {
	ast := assert.New(t)
	l := New()

	locked := l.TryLockWithTimeout(time.Second)
	ast.True(locked)

	start := time.Now()
	locked = l.TryLockWithTimeout(time.Second)
	d := time.Since(start)
	ast.False(locked)

	ast.Greater(time.Millisecond*100, d-time.Second)
}
