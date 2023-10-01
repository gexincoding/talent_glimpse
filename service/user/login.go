package user

import "talent_glimpse/core/smtp"

func SendEmailVerificationCode(code, to string) {
	smtp.SendEmail([]string{to}, "vertifucode", code)

}
