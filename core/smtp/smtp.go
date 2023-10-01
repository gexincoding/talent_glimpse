package smtp

import (
	"net/smtp"
	"talent_glimpse/core/config"

	"github.com/jordan-wright/email"
)

func SendEmail(to []string, subject, message string) error {
	e := email.NewEmail()

	//设置发送方的邮箱
	e.From = config.TalentGlimpseConfig.EmailConfig.From
	addr := config.TalentGlimpseConfig.EmailConfig.Addr
	username := config.TalentGlimpseConfig.EmailConfig.Username
	pass := config.TalentGlimpseConfig.EmailConfig.Pass
	host := config.TalentGlimpseConfig.EmailConfig.Host
	// 设置接收方的邮箱
	//	e.To = []string{"gexincoding@gmail.com"}
	e.To = to
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(message)

	//设置服务器相关的配置
	err := e.Send(addr, smtp.PlainAuth("", username, pass, host))
	if err != nil {
		return err
	}
	return nil
}
