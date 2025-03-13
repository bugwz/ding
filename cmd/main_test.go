package main

import (
	"os"
	"testing"

	"github.com/bugwz/ding/pkg"
)

func TestLoadConfig(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	configPath := home + "/.dingconf"
	_, err = os.Stat(configPath)

	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if os.IsNotExist(err) {
		if config != defaultConfig {
			t.Errorf("Expected default config, got %s", config)
		}
	} else {
		// 可以添加更多的配置文件存在时的测试逻辑
	}
}

func TestMainFunction(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// 测试无参数的情况
	os.Args = []string{"ding"}
	main()

	// 测试有参数的情况
	os.Args = []string{"ding", "mail,sms"}
	main()
}

func TestGetNotifierAndSendNotification(t *testing.T) {
	args := []string{"mail", "sms"}
	for _, arg := range args {
		notifier, err := pkg.GetNotifier(arg)
		if err != nil {
			t.Fatalf("Failed to get notifier for %s: %v", arg, err)
		}

		err = notifier.SendNotification("recipient@example.com", "Test Subject", "Test Body")
		if err != nil {
			t.Errorf("Failed to send notification using %s: %v", arg, err)
		}
	}
}
