package user

import "talent_glimpse/core/smtp"

var (
	CertificationSent = "[TalentGlimpse] 请您查收 talent glimpse 验证码，请勿泄露"
)

func SendEmailVerificationCode(code, to string) error {
	err := smtp.SendEmail([]string{to}, CertificationSent, code)
	if err != nil {
		return err
	}
	return nil
}
