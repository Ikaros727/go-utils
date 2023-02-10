package lock

import (
	"github.com/Icarus-0727/go-utils/pkg/lock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLock_TryLock(t *testing.T) {
	ast := assert.New(t)
	l := lock.New()

	locked := l.TryLock()
	ast.True(locked)

	locked = l.TryLock()
	ast.False(locked)
}

func TestLock_TimeoutLock(t *testing.T) {
	ast := assert.New(t)
	l := lock.New()

	locked := l.TimeoutLock(time.Second)
	ast.True(locked)

	start := time.Now()
	locked = l.TimeoutLock(time.Second)
	d := time.Since(start)
	ast.False(locked)

	ast.Greater(time.Millisecond*100, d-time.Second)
}
