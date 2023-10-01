package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"comment:用户邮箱;type:varchar(100)"`
	Phone    string `gorm:"comment:用户手机号;type:varchar(11)"`
	UserName string `gorm:"comment:用户名;type:varchar(100)"`
	Password string `gorm:"comment:密码;type:varchar(1000)"`
	ImageURL string `gorm:"comment:头像;type:varchar(1000)"`
}

func (User) TableName() string {
	return "user"
}
