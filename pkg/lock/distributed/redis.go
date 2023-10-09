package lock_distributed

import (
	"context"
	"time"

	"github.com/Icarus-0727/go-utils/pkg/lock"
	"github.com/go-redis/redis"
)

// 基于 Redis 的分布式锁
// client Redis 客户端
// lockKey Redis 用于标记上锁的 Key
// expired Redis 用于标记上锁的 Key 的过期时间
// localLock 基于 channel 的本地锁，用于实现单应用中锁的互斥，降低对 Redis 的操作
type distributedLockRedis struct {
	client    *redis.Client
	lockKey   string
	expired   time.Duration
	localLock lock.Lock
}

func (d *distributedLockRedis) Lock() {
	// lock in localLock
	d.localLock.Lock()

	// lock in Redis
	for {
		success, err := d.client.SetNX(d.lockKey, 1, d.expired).Result()
		// lock success, return
		if err == nil && success {
			return
		}
	}
}

func (d *distributedLockRedis) TryLock() (locked bool) {
	// try lock for localLock
	if d.localLock.TryLock() {
		// try lock for Redis
		success, err := d.client.SetNX(d.lockKey, 1, d.expired).Result()
		// try lock for Redis success, set locked = true
		if err == nil && success {
			locked = true
		} else {
			// try lock for Redis failed, unlock for localLock
			d.localLock.Unlock()
		}
	}

	return
}

func (d *distributedLockRedis) TryLockWithTimeout(timeout time.Duration) (locked bool) {
	// Start time, which is used to assist in calc the remaining timeout
	startTime := time.Now()

	// try lock for localLock with timeout
	if d.localLock.TryLockWithTimeout(timeout) {
		// calc the remaining timeout
		timeout -= time.Since(startTime) + time.Millisecond
		// create timeout ctx
		timeoutCtx, cancelFunc := context.WithTimeout(context.TODO(), timeout)
		defer cancelFunc()

		for {
			select {
			// timeout
			case <-timeoutCtx.Done():
				d.localLock.Unlock()

			// try lock for Redis
			default:
				success, err := d.client.SetNX(d.lockKey, 1, d.expired).Result()
				// try lock for Redis success, return
				if err == nil && success {
					locked = true
					return
				}
			}
		}
	}

	return
}

func (d *distributedLockRedis) Unlock() {
	// unlock for localLock
	defer d.localLock.Unlock()
	// unlock for Redis
	d.client.Del(d.lockKey)
}

// NewDistributedLockRedis 初始化
func NewDistributedLockRedis(client *redis.Client, lockKey string, expired time.Duration) lock.Lock {
	return &distributedLockRedis{
		client:    client,
		lockKey:   lockKey,
		expired:   expired,
		localLock: lock.New(),
	}
}
