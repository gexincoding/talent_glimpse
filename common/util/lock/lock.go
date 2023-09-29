package lock

import (
	"log"
	"talent_glimpse/common/rdb"
	"time"
)

/*

// Redis is rdb client
var (
	Base                       = "fcdn_bizapi_"
	BaseICP                    = "fcdn_bizapi_icp_"
	GlobalUpdateLockKey        = Base + "global_update_lock"
	GlobalUpdateLockMonitorKey = Base + "global_update_lock_monitor"
	GlobalCertAlarmLockKey     = Base + "global_cert_alarm"
	GlobalCertUpdateLockKey    = Base + "global_cert_update_lock"
	GlobalUpdateLockTTL        = time.Second * 3

	historyPrefix = "stream_history_info_"

	MonitorKeysTTL  = time.Minute * 3
	MonitorKeyAll   = Base + "monitor:all"
	MonitorKeyIn    = Base + "monitor:in"
	MonitorKeyNotIn = Base + "monitor:not_in"

	CertCacheKey = Base + "cert:"
	CertJudgeKey = Base + "cert_judge:"
)*/

/*// Initialize rdb client according to config
func InitRedis(conf config.RedisConfig) error {
	if conf.Psm == "" {
		return nil
	}
	var err error
	option := goredis.NewOption()
	option.MaxRetries = 10 // retry times
	if env.IsProduct() || len(conf.Hosts) == 0 {
		Redis, err = goredis.NewClientWithOption(conf.Psm, option)
	} else {
		option.DisableAutoLoadConf()
		Redis, err = goredis.NewClientWithServers(conf.Psm, conf.Hosts, option)
	}
	if err != nil {
		return errors.Wrap(err, "new client error")
	}

	return nil
}
*/
// Acquire a lock by SetNX
func AcquireLock(key string, ttl time.Duration, timeout ...time.Duration) bool {
	endTime := time.Now()
	if len(timeout) != 0 && timeout[0] != 0 {
		endTime = endTime.Add(timeout[0])
	}

	acquireLock := func() bool {
		ok, err := rdb.SetNX(key, "value", ttl)
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

/*
// get all the vhosts which have been pushed stream in 30 days
func GetAllVhosts(ctx context.Context, user string) (allVhosts []string, err error) {
	key := Base + historyPrefix + "_all_vhost_" + user
	return GetAllElementsInSomeDays(ctx, key, int64(30))
}
*/
/*func GetAllPushDomains(ctx context.Context, vhost string) (pushDomains []string, err error) {
	key := Base + historyPrefix + "all_push_domain_" + vhost
	return GetAllElementsInSomeDays(ctx, key, int64(30))
}*/

/*func GetAllApps(ctx context.Context, pushDomain string) (apps []string, err error) {
	key := Base + historyPrefix + "all_app_" + pushDomain
	return GetAllElementsInSomeDays(ctx, key, int64(30))
}*/
/*
// 从sorted set里面获取最近30天内的信息， 是个通用封装，用于查询vhost、domain、app信息
// 使用的基本逻辑是：
// 找出最近30天推过流的所有vhost
// 对于每一个vhost，找出最近30天推过流的所有push domain
// 对于每一个push domain， 找出最近30天推过流的所有app
func GetAllElementsInSomeDays(ctx context.Context, key string, days int64) (allVhosts []string, err error) {
	ts := time.Now().Unix()
	timestampStop := strconv.FormatInt(ts, 10)
	timestampStart := strconv.FormatInt(ts-86400*days, 10)
	cmd := Redis.ZRangeByScore(key, redis.ZRangeBy{
		Min:    timestampStart,
		Max:    timestampStop,
		Offset: 0,
		Count:  -1,
	})
	return cmd.Result()
}
*/
