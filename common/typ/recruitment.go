package typ

import (
	"time"
)

type CreateRecruitmentInfoReq struct {
	RecruitmentYear      int       `json:"RecruitmentYear"`
	RecruitmentWebsite   string    `json:"RecruitmentWebsite"`
	WrittenTestStartTime time.Time `json:"WrittenTestStartTime"`
	WrittenTestEndTime   time.Time `json:"WrittenTestEndTime"`
	InterviewStartTime   time.Time `json:"InterviewStartTime"`
	InterviewEndTime     time.Time `json:"InterviewEndTime"`
	RecruitmentStartTime time.Time `json:"RecruitmentStartTime"`
	RecruitmentEndTime   time.Time `json:"RecruitmentEndTime"`
	FrontendHaving       bool      `json:"FrontendHaving"`
	BackendHaving        bool      `json:"BackendHaving"`
	AlgorithmHaving      bool      `json:"AlgorithmHaving"`
	HardwareHaving       bool      `json:"HardwareHaving"`
}
