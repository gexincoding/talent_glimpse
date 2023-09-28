package redis

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"talent_glimpse/core/config"
)

var Rdb *redis.Client

func Init(conf config.RedisConfig) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Pass,
	})
	if Rdb == nil {
		err := errors.New("init redis err")
		log.Println(fmt.Errorf("init redis error: %v", err))
		return err
	}
	return nil
}
