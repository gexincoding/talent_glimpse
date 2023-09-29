package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"talent_glimpse/core/config"
	"talent_glimpse/core/database"
	"talent_glimpse/core/redis"
	"talent_glimpse/handlers"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
	Age   int
}

func Test() {
	//redis.Client.Set(context.Background(), "kkkk", "vvvv", 20*time.Second)
	////database.DB.AutoMigrate(&User{})
	//user := User{
	//	Name:  "John Doe",
	//	Email: "johndoe@example.com",
	//	Age:   25,
	//}
	//database.DB.Create(&user)
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
		recruitmentRouter := router.Group("recruitment")
		recruitmentRouter.POST("/CreateRecruitmentInfo", handlers.CreateRecruitmentInfo)
	}

	{

	}

	{

	}

	{

	}
	Test()

	err = r.Run(":23462")
	if err != nil {
		return
	}

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
