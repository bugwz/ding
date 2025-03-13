package pkg

import (
	"fmt"
	"log"
	"os/exec"
)

// SystemConfig 系统提醒配置
type SystemConfig struct {
	// 可添加具体配置字段
}

// SendSystemNotification 发送系统弹窗提醒
func SendSystemNotification(message string) error {
	// 调用 macOS 系统的 osascript 命令来显示弹窗提醒
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%s\" with title \"系统提醒\"", message))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("执行系统提醒命令失败: %v, 输出: %s", err, string(output))
		return err
	}
	return nil
}
