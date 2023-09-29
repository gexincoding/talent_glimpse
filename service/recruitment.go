package service

import (
	"talent_glimpse/common/db"
	"talent_glimpse/common/db/model"
	"talent_glimpse/common/typ"
)

func CreateRecruitmentInfo(req typ.CreateRecruitmentInfoReq) error {
	info := model.RecruitmentInfo{
		Year:                 req.Year,
		Targets:              req.Targets,
		Website:              req.Website,
		CompanyID:            req.CompanyID,
		WrittenTestStartTime: req.WrittenTestStartTime,
		WrittenTestEndTime:   req.WrittenTestEndTime,
		InterviewStartTime:   req.InterviewStartTime,
		InterviewEndTime:     req.InterviewEndTime,
		RecruitmentStartTime: req.RecruitmentStartTime,
		RecruitmentEndTime:   req.RecruitmentEndTime,
		FrontendHaving:       req.FrontendHaving,
		BackendHaving:        req.BackendHaving,
		AlgorithmHaving:      req.AlgorithmHaving,
		HardwareHaving:       req.HardwareHaving,
		ProductsMgrHaving:    req.ProductsMgrHaving,
		Status:               0,
	}
	err := db.CreateRecruitmentInfo(info)
	if err != nil {
		return err
	}
	return nil
}
