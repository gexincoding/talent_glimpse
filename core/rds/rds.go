package rds

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"talent_glimpse/core/config"
)

var Client *redis.Client

func Init(conf config.RedisConfig) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Pass,
	})
	if Client == nil {
		err := errors.New("init rdb err")
		log.Println(fmt.Errorf("init rdb error: %v", err))
		return err
	}
	return nil
}
