package user

import (
	"talent_glimpse/core/config"
	"testing"
)

func TestSendEmailVerificationCode(t *testing.T) {
	_ = config.Init("../../conf/conf.json")
	SendEmailVerificationCode("837413", "gexincoding@gmail.com")
}
