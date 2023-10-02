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
	rdb.Set("testtesttest", "11111111", time.Second*100)
	//rdb.HMSet("test", map[string]string{"11": "222"})
	//rdb.RefreshList("refresh", []interface{}{"a", "b", "c"})
}

func main() {
	// 初始化：连接需要用到的一些数据库、网络组件
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
		useRrouter := router.Group("/user")
		useRrouter.POST("/Login", handlers.Login)
		useRrouter.POST("/Logout", handlers.Logout)
		useRrouter.POST("/ChangePassword", handlers.ChangePassword)
		useRrouter.GET("/GetUserDetail", handlers.GetUserDetail)
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
	// 把配置文件里的内容读到当前进程的某个结构体内
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
