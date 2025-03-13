package notify

import (
	"fmt"

	"github.com/twilio/twilio-go"
)

// SMSConfig 短信配置
type SMSConfig struct {
	Provider   string
	AccountSID string
	AuthToken  string
	From       string
}

// SendSMS 发送短信
func SendSMS(config SMSConfig, to string, message string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AccountSID,
		Password: config.AuthToken,
	})

	params := &v2010.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(config.From)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	return err
}

// SendMultipleNotifications 发送多种通知
func SendMultipleNotifications(configs map[string]interface{}, recipients map[string]string, message string) error {
	for notifType, config := range configs {
		switch notifType {
		case "mail":
			mailConfig, ok := config.(MailConfig)
			if !ok {
				return fmt.Errorf("invalid mail config")
			}
			err := SendMail(mailConfig, recipients[notifType], "Test Subject", message)
			if err != nil {
				return err
			}
		case "sms":
			smsConfig, ok := config.(SMSConfig)
			if !ok {
				return fmt.Errorf("invalid sms config")
			}
			err := SendSMS(smsConfig, recipients[notifType], message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
