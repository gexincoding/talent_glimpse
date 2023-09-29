package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name           string `gorm:"comment:公司名称;type:varchar(100)"`
	CompanyTypeID  uint   `gorm:"comment:公司类型ID"`
	Description    string `gorm:"comment:公司简介"`
	Website        string `gorm:"comment:公司官网;type:varchar(1000)"`
	CareersWebsite string `gorm:"comment:招聘官网;type:varchar(1000)"`
	Icon           string `gorm:"comment:公司图标;type:varchar(1000)"`
	Status         int    `gorm:"comment:状态，0未删除，1已删除"`
}

type CompanyType struct {
	gorm.Model
	Name   string `gorm:"comment:公司类型名;type:varchar(20)"`
	Status int    `gorm:"comment:状态，0未删除，1已删除"`
}

func (CompanyType) TableName() string {
	return "company_type"
}

func (Company) TableName() string {
	return "company"
}
