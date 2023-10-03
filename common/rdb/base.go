package rdb

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"talent_glimpse/core/rds"
	"time"
)

const MaxStep = 50

var MonitorKeysTTL = time.Minute * 3

func HSet(key string, value map[string]interface{}) (int64, error) {
	return rds.Client.HSet(context.Background(), key, value).Result()

}

func HGet(key, field string) (string, error) {
	return rds.Client.HGet(context.Background(), key, field).Result()
}

// SetNX wrap rdb SetNX
func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return rds.Client.SetNX(context.Background(), key, value, expiration).Result()
}

// Del wrap rdb Del
func Del(keys ...string) (int64, error) {
	return rds.Client.Del(context.Background(), keys...).Result()
}

// Expire wrap rdb Expire
func Expire(key string, expiration time.Duration) (bool, error) {
	return rds.Client.Expire(context.Background(), key, expiration).Result()
}

// ExpireAt wrap rdb ExpireAt
func ExpireAt(key string, time time.Time) (bool, error) {
	return rds.Client.ExpireAt(context.Background(), key, time).Result()
}

// Set wrap rdb Set
func Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return rds.Client.Set(context.Background(), key, value, expiration).Result()
}

// Get wrap rdb Get
func Get(key string) (string, error) {
	return rds.Client.Get(context.Background(), key).Result()
}

func HMSet(key string, value map[string]string) error {
	count := 0
	tempMap := make(map[string]interface{})
	for k, v := range value {
		count++
		tempMap[k] = v
		if count == MaxStep {
			_, err := rds.Client.HMSet(context.Background(), key, tempMap).Result()
			if err != nil {
				return err
			}
			count = 0
			tempMap = make(map[string]interface{})
		}
	}
	if count > 0 {
		_, err := rds.Client.HMSet(context.Background(), key, tempMap).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func HGetAll(key string) (map[string]string, error) {
	return rds.Client.HGetAll(context.Background(), key).Result()
}

func HDel(key string, fields ...string) error {
	if len(fields) > MaxStep { // 单次操作key数量设置上限，防止redis阻塞
		var i int
		for i = 0; i < len(fields)/MaxStep; i++ {
			_, err := rds.Client.HDel(context.Background(), key, fields[i*MaxStep:(i+1)*MaxStep]...).Result()
			if err != nil {
				return err
			}
		}
		_, err := rds.Client.HDel(context.Background(), key, fields[i*MaxStep:]...).Result()
		if err != nil {
			return err
		}
	} else {
		_, err := rds.Client.HDel(context.Background(), key, fields...).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func RefreshHashMap(key string, value map[string]string) error {
	originData, err := HGetAll(key)
	if err != nil {
		return errors.New("get origin data error when refresh hash")
	}
	diffKeys := make([]string, 0) // keys exists in origin but not in new map
	for originKey := range originData {
		if _, ok := value[originKey]; !ok {
			diffKeys = append(diffKeys, originKey)
		}
	}
	err = HDel(key, diffKeys...)
	if err != nil {
		return errors.New("del origin key error when refresh hash")
	}
	err = HMSet(key, value)
	if err != nil {
		return errors.New("HMSet error when refresh hash")
	}
	_, err = Expire(key, MonitorKeysTTL)
	if err != nil {
		return err
	}
	return nil
}

func pipelineGet(getList []string) ([]interface{}, error) {
	ctx := context.Background()
	if len(getList) == 0 {
		return []interface{}{}, nil
	}
	pipe := rds.Client.Pipeline()
	//	defer func() { _ = pipe.Close() }()
	cmds := make([]*redis.StringCmd, len(getList))
	result := make([]interface{}, len(getList))
	for i, key := range getList {
		cmds[i] = pipe.Get(ctx, key)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return result, err
	}
	for i, cmd := range cmds {
		value, err := cmd.Result()
		if err != nil {
			result[i] = nil
		} else {
			result[i] = value
		}
	}
	return result, nil
}

func Mget(keys []string, pageSize int) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	step := pageSize
	if len(keys) > step { // 单次操作key数量设置上限，防止redis阻塞
		var i int
		for i = 0; i < len(keys)/step; i++ {
			values, err := pipelineGet(keys[i*step : (i+1)*step])
			if err != nil {
				return nil, err
			}
			ret = append(ret, values...)
		}
		values, err := pipelineGet(keys[i*step:])
		if err != nil {
			return nil, err
		}
		ret = append(ret, values...)

	} else {
		values, err := pipelineGet(keys)
		if err != nil {
			return nil, err
		}
		ret = append(ret, values...)
	}
	return ret, nil
}

func GetListAll(key string) ([]string, error) {
	return rds.Client.LRange(context.Background(), key, 0, -1).Result()
}

func RefreshList(key string, elems []interface{}) error {
	ctx := context.Background()
	pipe := rds.Client.Pipeline()
	//defer func() { _ = pipe.Close() }()
	pipe.Del(ctx, key)
	pipe.LPush(ctx, key, elems...)
	pipe.Expire(ctx, key, MonitorKeysTTL)
	_, err := pipe.Exec(ctx)
	return err
}
