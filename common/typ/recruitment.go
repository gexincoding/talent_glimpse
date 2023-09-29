package typ

import (
	"time"
)

type CreateRecruitmentInfoReq struct {
	Year                 int
	Targets              string
	Website              string
	CompanyID            uint
	WrittenTestStartTime *time.Time
	WrittenTestEndTime   *time.Time
	InterviewStartTime   *time.Time
	InterviewEndTime     *time.Time
	RecruitmentStartTime *time.Time
	RecruitmentEndTime   *time.Time
	FrontendHaving       *bool
	BackendHaving        *bool
	AlgorithmHaving      *bool
	HardwareHaving       *bool
	ProductsMgrHaving    *bool
}
