package main

import (
	"context"
	"fmt"
	"log"
	"talent_glimpse/core/config"
	"talent_glimpse/core/database"
	"talent_glimpse/core/redis"
	"time"
)

func main() {
	err := Init()
	if err != nil {
		log.Println(fmt.Errorf("init error: %v", err))
	}
	redis.Rdb.Set(context.Background(), "k1", "v1", 100*time.Second)
}

func Init() error {
	err := config.Init()
	if err != nil {
		log.Println(fmt.Errorf("config init error: %v", err))
		return err
	}
	err = database.Init(config.TalentGlimpseConfig.MySQLConfig)
	if err != nil {
		log.Println(fmt.Errorf("database init error: %v", err))
		return err
	}
	err = redis.Init(config.TalentGlimpseConfig.RedisConfig)
	if err != nil {
		log.Println(fmt.Errorf("redis init error: %v", err))
		return err
	}
	return nil

}
