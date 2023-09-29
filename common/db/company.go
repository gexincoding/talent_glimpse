package db

import (
	"fmt"
	"log"
	"talent_glimpse/common/db/model"
	"talent_glimpse/core/database"
)

func CreateRecruitmentInfo(info model.RecruitmentInfo) error {
	res := database.Client.Create(&info)
	if res.Error != nil || res.RowsAffected < 1 {
		log.Println(fmt.Errorf("[CreateRecruitmentInfo] db insert wrong"))
	}
	return nil
}
