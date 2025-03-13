package pkg

import (
	"fmt"
	"net/smtp"
)

// MailConfig 邮件配置
type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// SendMail 发送邮件
func SendMail(config MailConfig, to string, subject string, body string) error {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		auth,
		config.From,
		[]string{to},
		[]byte(msg),
	)
}
