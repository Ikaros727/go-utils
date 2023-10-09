package lock_distributed

import (
	"testing"
	"time"

	lock_distributed "github.com/Icarus-0727/go-utils/pkg/lock/distributed"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func Test_distributedLockRedis_TryLock(t *testing.T) {
	ast := assert.New(t)
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: s.Addr()})

	lock := lock_distributed.NewDistributedLockRedis(client, "unittest", 0)
	lock2 := lock_distributed.NewDistributedLockRedis(client, "unittest", 0)

	// 都没锁，加锁
	ast.True(lock.TryLock())
	lock.Unlock()
	ast.True(lock2.TryLock())
	lock2.Unlock()

	// 已上锁，加锁
	lock.TryLock()
	ast.False(lock.TryLock())
	ast.False(lock2.TryLock())
}

func Test_distributedLockRedis_TryLockWithTimeout(t *testing.T) {
	ast := assert.New(t)
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: s.Addr()})

	lock := lock_distributed.NewDistributedLockRedis(client, "unittest", time.Second*5)
	lock2 := lock_distributed.NewDistributedLockRedis(client, "unittest", time.Second*5)

	// 都没锁，加锁
	ast.True(lock.TryLockWithTimeout(time.Second * 5))
	lock.Unlock()
	ast.True(lock2.TryLockWithTimeout(time.Second * 5))
	lock2.Unlock()

	// 已上锁，加锁
	lock.Lock()
	//ast.False(lock.TryLockWithTimeout(time.Second * 5))
	//ast.False(lock2.TryLockWithTimeout(time.Second * 5))
}
