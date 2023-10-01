package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"talent_glimpse/common/rdb"
	"talent_glimpse/common/util"
	"talent_glimpse/core/config"
	"talent_glimpse/core/database"
	"talent_glimpse/core/rds"
	"talent_glimpse/handlers"
	"time"
)

func Test() {
	rdb.Set("1", "11", time.Second*100)
	rdb.HMSet("test", map[string]string{"11": "222"})
	rdb.RefreshList("refresh", []interface{}{"a", "b", "c"})
}

func main() {

	err := Init()
	if err != nil {
		panic(errors.New("main init failed"))
	}
	r := gin.Default()

	router := r.Group("/talent_glimpse/v1")
	{
		testRouter := router.Group("")
		testRouter.GET("/Ping", handlers.Ping)
		testRouter.POST("/Ping", handlers.Ping)
	}

	{
		recruitmentRouter := router.Group("/recruitment")
		recruitmentRouter.POST("/CreateRecruitmentInfo", handlers.CreateRecruitmentInfo)
	}

	Test()

	err = r.Run(":23462")
	if err != nil {
		return
	}

}

func Init() error {
	// 当前进程的ID
	util.ProcessID = util.NewID()
	err := config.Init("conf/conf.json")
	if err != nil {
		log.Println(fmt.Errorf("config init error: %v", err))
		return err
	}
	err = database.Init(config.TalentGlimpseConfig.MySQLConfig)
	if err != nil {
		log.Println(fmt.Errorf("database init error: %v", err))
		return err
	}
	err = rds.Init(config.TalentGlimpseConfig.RedisConfig)
	if err != nil {
		log.Println(fmt.Errorf("rdb init error: %v", err))
		return err
	}
	return nil

}
