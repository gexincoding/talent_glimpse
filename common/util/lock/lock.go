package lock

import (
	"log"
	"talent_glimpse/common/rdb"
	"time"
)

// Acquire a lock by SetNX
func AcquireLock(key string, ttl time.Duration, timeout ...time.Duration) bool {
	endTime := time.Now()
	if len(timeout) != 0 && timeout[0] != 0 {
		endTime = endTime.Add(timeout[0])
	}

	acquireLock := func() bool {
		ok, err := rdb.SetNX(key, "test", ttl)
		if err != nil || !ok {
			//logs.Warn("acquire lock error, err: %v key: %s ttl: %v", err, key, ttl)
			return false
		}
		return true
	}

	var isLock bool
	for {
		if isLock = acquireLock(); isLock {
			return isLock
		}
		if endTime.Before(time.Now()) {
			return false
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// Release a lock
func ReleaseLock(key string) {
	if n, err := rdb.Del(key); err != nil {
		log.Printf("[Warn] release lock error, err: %s key: %s", err, key)
	} else if n == 0 {
		log.Printf("[Warn] lock already released, key: %s", key)
	}
}

// ExpireLock reset lock ttl
func ExpireLock(key string, ttl time.Duration) {
	if ok, err := rdb.Expire(key, ttl); err != nil {
		log.Printf("[Warn] expire lock time error, err: %s key: %s ttl: %d", err, key, ttl)
	} else if !ok {
		log.Printf("[Warn] lock already expired, key: %s", key)
	}
}

// ExpireLock reset lock ttl
func ExpireAtLock(key string, time time.Time) bool {
	ok, err := rdb.ExpireAt(key, time)
	if err != nil {
		log.Printf("[Warn] expire lock time error, err: %s key: %s time: %d", err, key, time)
		return false
	} else if !ok {
		log.Printf("[Warn] lock already expired, key: %s", key)
	}
	return ok
}
