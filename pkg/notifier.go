package pkg

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// Notifier notification interface
type Notifier interface {
	SendNotification(recipient, subject, message string) error
}

// MailNotifier email notification structure
type MailNotifier struct {
	Config MailConfig
}

// SendNotification implements the email notification interface method
func (m *MailNotifier) SendNotification(recipient, subject, message string) error {
	return SendMail(m.Config, recipient, subject, message)
}

// SMSNotifier SMS notification structure
type SMSNotifier struct {
	Config SMSConfig
}

// SendNotification implements the SMS notification interface method
func (s *SMSNotifier) SendNotification(recipient, subject, message string) error {
	return SendSMS(s.Config, recipient, message)
}

// GetNotifier gets the notifier based on the configuration
func GetNotifier(configType string) (Notifier, error) {
	config, err := ini.Load("config/default.ini")
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %v", err)
	}
	switch configType {
	case "mail":
		mailConfig := MailConfig{
			Host:     config.Section("mail").Key("host").String(),
			Port:     config.Section("mail").Key("port").MustInt(),
			Username: config.Section("mail").Key("username").String(),
			Password: config.Section("mail").Key("password").String(),
			From:     config.Section("mail").Key("from").String(),
		}
		return &MailNotifier{Config: mailConfig}, nil
	case "sms":
		smsConfig := SMSConfig{
			Provider:   config.Section("sms").Key("provider").String(),
			AccountSID: config.Section("sms").Key("account_sid").String(),
			AuthToken:  config.Section("sms").Key("auth_token").String(),
			From:       config.Section("sms").Key("from").String(),
		}
		return &SMSNotifier{Config: smsConfig}, nil
	default:
		return nil, fmt.Errorf("Unsupported notification type: %s", configType)
	}
}
