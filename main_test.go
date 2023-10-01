package main

import (
	"talent_glimpse/common/db/model"
	"talent_glimpse/core/database"
	"testing"
)

func TestGenerate(t *testing.T) {
	_ = Init()
	migrator := database.Client.Migrator()
	_ = migrator.DropTable(&model.RecruitmentInfo{})
	_ = migrator.DropTable(&model.Company{})
	_ = migrator.DropTable(&model.RecruitmentComment{})
	_ = migrator.DropTable(&model.CommentLike{})
	_ = migrator.DropTable(&model.ReferralInfo{})
	_ = migrator.DropTable(&model.CareerDirection{})
	_ = migrator.DropTable(&model.CompanyType{})
	_ = migrator.DropTable(&model.User{})

	_ = migrator.CreateTable(&model.RecruitmentInfo{})
	_ = migrator.CreateTable(&model.Company{})
	_ = migrator.CreateTable(&model.RecruitmentComment{})
	_ = migrator.CreateTable(&model.CommentLike{})
	_ = migrator.CreateTable(&model.ReferralInfo{})
	_ = migrator.CreateTable(&model.CareerDirection{})
	_ = migrator.CreateTable(&model.CompanyType{})
	_ = migrator.CreateTable(&model.User{})

}
