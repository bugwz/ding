package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bugwz/ding/pkg"
)

var Version = "dev" // Version information is injected at compile time

// Default configuration
var defaultConfig = `{
	"setting1": "value1",
	"setting2": "value2"
}`

func loadConfig() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return defaultConfig, err
	}

	configPath := filepath.Join(home, ".dingconf")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return defaultConfig, nil
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return defaultConfig, err
	}

	return string(content), nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	fmt.Printf("Using config: %s\n", config)
	fmt.Printf("Ding Tool %s\n", Version)

	if len(os.Args) > 1 {
		// Support multiple parameters, e.g., ding mail,sms
		args := strings.Split(os.Args[1], ",")
		for _, arg := range args {
			notifier, err := pkg.GetNotifier(arg)
			if err != nil {
				fmt.Printf("Error getting notifier: %v\n", err)
				continue
			}
			err = notifier.SendNotification("recipient@example.com", "Test Subject", "Test Body")
			if err != nil {
				fmt.Printf("Failed to send notification: %v\n", err)
			}
		}
	} else {
		fmt.Println("Usage: ding [mail|sms|mail,sms]")
	}
}
