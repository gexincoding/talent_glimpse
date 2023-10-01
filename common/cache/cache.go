package cache

import (
	"fmt"
	"talent_glimpse/common/rdb"
	"time"
)

func Put(cacheName, k string, v interface{}, expireTime time.Duration) error {
	_, err := rdb.Set(fmt.Sprintf("%s_%s", cacheName, k), v, expireTime)
	if err != nil {
		return err
	}
	return nil
}

func Get(cacheName, k string) (interface{}, bool) {
	v, err := rdb.Get(fmt.Sprintf("%s_%s", cacheName, k))
	if err != nil {
		return nil, false
	}
	return v, true
}
