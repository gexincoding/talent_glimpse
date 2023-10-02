package handlers

import (
	"github.com/gin-gonic/gin"
	"talent_glimpse/common/cache"
	"talent_glimpse/common/db/model"
	"time"
)

var (
	TokenExpireTime = 30 * 24 * 60 * time.Minute
)

func Login(ctx *gin.Context) {
	// 1. 校验邮箱和验证码是否正确
	//   1.1 校验邮箱格式、redis中是否有这个邮箱的验证码记录
	//   1.2 去redis查看邮箱的验证码是否匹配，无论是否匹配，删除redis的key
	// 2. 数据库有没有这个邮箱，如果没有，插入新的
	// 3. 生成token
	// 4. 将token和用户信息放在redis中
	// 5. 返回token

	token := "1r42tfdw142ytwrgfewd12f3546jyhtbrvwefd123435ythgref"
	//email := ctx.Query("Email")
	//code := ctx.Query("Code")
	err := cache.Put(cache.LoginUserCache, token, &model.User{
		Email:    "",
		Phone:    "",
		UserName: "",
		Password: "",
		ImageURL: "",
	}, TokenExpireTime)
	if err != nil {
		return
	}

}

func Logout(ctx *gin.Context) {

}

func ChangePassword(ctx *gin.Context) {

}

func GetUserDetail(ctx *gin.Context) {

}
