package user

import (
	"fmt"
	"talent_glimpse/core/smtp"
)

var (
	CertificationSent = "[TalentGlimpse] 请您查收 talent glimpse 验证码，请勿泄露"
)

func SendEmailVerificationCode(code, to string) {
	msg := fmt.Sprintf("尊敬的用户：\n\n        您好！非常感谢您使用TalentGlimpse，您正在尝试使用邮箱登录。\n\n        您的验证码是：%s\n\n        如果此并非您的操作，请忽视，切勿将验证码泄露给他人。", code)
	_ = smtp.SendEmail([]string{to}, CertificationSent, msg)
}
