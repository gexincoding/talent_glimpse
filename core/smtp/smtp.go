package smtp

import (
	"net/smtp"
	"talent_glimpse/core/config"

	"github.com/jordan-wright/email"
)

func SendEmail(to []string, subject, message string) error {
	e := email.NewEmail()

	e.To = to
	e.Subject = subject
	e.Text = []byte(message)
	e.From = config.TalentGlimpseConfig.EmailConfig.From
	addr := config.TalentGlimpseConfig.EmailConfig.Addr
	username := config.TalentGlimpseConfig.EmailConfig.Username
	pass := config.TalentGlimpseConfig.EmailConfig.Pass
	host := config.TalentGlimpseConfig.EmailConfig.Host

	//设置服务器相关的配置
	err := e.Send(addr, smtp.PlainAuth("", username, pass, host))
	if err != nil {
		return err
	}

	return nil
}
