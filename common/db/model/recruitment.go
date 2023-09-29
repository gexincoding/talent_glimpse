package model

import (
	"gorm.io/gorm"
	"time"
)

// RecruitmentComment 招聘信息评论表
type RecruitmentComment struct {
	gorm.Model
	Content         string    `gorm:"comment:评论内容"`
	ParentCommentID uint      `gorm:"comment:父级评论ID"`
	UserID          uint      `gorm:"comment:评论用户ID"`
	CommentTime     time.Time `gorm:"comment:评论时间"`
	Status          int       `gorm:"comment:状态，0未删除，1已删除"`
}

type CommentLike struct {
	gorm.Model
	UserID    uint `gorm:"comment:点赞的用户id"`
	CommentID uint `gorm:"comment:评论ID"`
	Status    int  `gorm:"comment:状态，0未删除，1已删除"`
}

type RecruitmentInfo struct {
	gorm.Model
	Year                 int        `gorm:"comment:招聘年份"`
	Targets              string     `gorm:"comment:招聘对象;type:varchar(100)"`
	Website              string     `gorm:"comment:招聘网站;type:varchar(1024)"`
	CompanyID            uint       `gorm:"comment:招聘的公司ID"`
	WrittenTestStartTime *time.Time `gorm:"comment:笔试开始时间"`
	WrittenTestEndTime   *time.Time `gorm:"comment:笔试结束时间"`
	InterviewStartTime   *time.Time `gorm:"comment:面试开始时间"`
	InterviewEndTime     *time.Time `gorm:"comment:面试结束时间"`
	RecruitmentStartTime *time.Time `gorm:"comment:招聘开始时间"`
	RecruitmentEndTime   *time.Time `gorm:"comment:招聘结束时间"`
	FrontendHaving       *bool      `gorm:"comment:是否有后端岗位"`
	BackendHaving        *bool      `gorm:"comment:是否有前端岗位"`
	AlgorithmHaving      *bool      `gorm:"comment:是否有算法岗位"`
	HardwareHaving       *bool      `gorm:"comment:是否有硬件岗位"`
	ProductsMgrHaving    *bool      `gorm:"comment:是否有产品/运营岗位"`
	Status               int        `gorm:"comment:状态，0未删除，1已删除"`
}

type ReferralInfo struct {
	gorm.Model
	RecruitmentInfoID uint    `gorm:"comment:招聘信息ID"`
	CompanyID         uint    `gorm:"comment:内推公司ID"`
	Contact           string  `gorm:"comment:内推人联系方式;type:varchar(100)"`
	ReferralCode      *string `gorm:"comment:内推码;type:varchar(20)"`
	ReferralLink      *string `gorm:"comment:内推链接;type:varchar(1000)"`
	PositionOrDept    *string `gorm:"comment:内推岗位/部门,type;varchar(100)"`
	Status            int     `gorm:"comment:状态，0未删除，1已删除"`
}
type CareerDirection struct {
	gorm.Model
	Name         string `gorm:"comment:名称;type:varchar(20)"`
	Introduction string `gorm:"comment:简介;type:varchar(1000)"`
	Status       int    `gorm:"comment:状态，0未删除，1已删除"`
}

// 设置表名为 `career_direction`
func (CareerDirection) TableName() string {
	return "career_direction"
}

// 设置表名为 `referral_info`
func (ReferralInfo) TableName() string {
	return "referral_info"
}

func (RecruitmentComment) TableName() string {
	return "recruitment_comment"
}
func (CommentLike) TableName() string {
	return "comment_like"
}

func (RecruitmentInfo) TableName() string {
	return "recruitment_info"
}
