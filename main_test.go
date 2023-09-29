package main

import (
	"gorm.io/gorm"
	"talent_glimpse/core/database"
	"testing"
	"time"
)

type RecruitmentInfo struct {
	gorm.Model
	// 招聘年份
	Year int
	//招聘对象
	Targets string
	// 招聘网站
	Website string
	// 招聘的公司ID
	CompanyID uint
	//笔试开始时间
	WrittenTestStartTime time.Time
	// 笔试结束时间
	WrittenTestEndTime time.Time
	// 面试开始时间
	InterviewStartTime time.Time
	// 面试结束时间
	InterviewEndTime time.Time
	// 招聘开始时间
	RecruitmentStartTime time.Time
	// 招聘结束时间
	RecruitmentEndTime time.Time
	// 是否有后端岗位
	FrontendHaving bool
	// 是否有前端岗位
	BackendHaving bool
	// 是否有算法岗位
	AlgorithmHaving bool
	// 是否有硬件岗位
	HardwareHaving bool
	// 是否有产品/运营岗位
	ProductsMgrHaving bool
}

type Company struct {
	gorm.Model
	// 公司名称
	Name string
	// 公司类型
	Type string
	// 公司简介
	Description string
	// 公司官网
	Website string
	//招聘官网
	CareersWebsite string
}

// RecruitmentComment 招聘信息评论表
type RecruitmentComment struct {
	gorm.Model
	Content         string    `gorm:"type:text;not null;comment:评论内容"`
	ParentCommentID uint      `gorm:"comment:父级评论ID"`
	UserID          uint      `gorm:"comment:评论用户ID"`
	CommentTime     time.Time `gorm:"comment:评论时间"`
}

func (RecruitmentComment) TableName() string {
	return "recruitment_comment"
}

func (Company) TableName() string {
	return "company"
}

func (RecruitmentInfo) TableName() string {
	return "recruitment_info"
}

func TestGenerate(t *testing.T) {
	_ = Init()
	_ = database.Client.AutoMigrate(&RecruitmentInfo{})
	_ = database.Client.AutoMigrate(&Company{})
	_ = database.Client.AutoMigrate(&RecruitmentComment{})

}
