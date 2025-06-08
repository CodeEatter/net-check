package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	ipv6Url = "https://check-ipv6.sangmin.eu.org/check.json"
	ipv4Url = "https://check-ipv4.sangmin.eu.org/check.json"
	timeout = 5 * time.Second
)

type Messages struct {
	CheckTitle  string
	IPv6        string
	IPv4        string
	Success     string
	Error       string
	EndMessage  string
}

func getMessages() Messages {
	lang := getLang()

	switch lang {
	case "ko":
		return Messages{
			CheckTitle: "네트워크 확인",
			IPv6:       "IPv6",
			IPv4:       "IPv4",
			Success:    "정상",
			Error:      "오류",
			EndMessage: "초 후에 종료",
		}
	case "ja":
		return Messages{
			CheckTitle: "ネットチェック",
			IPv6:       "IPv6",
			IPv4:       "IPv4",
			Success:    "正常",
			Error:      "エラー",
			EndMessage: "秒後に終了",
		}
	case "zh":
		return Messages{
			CheckTitle: "网络检查",
			IPv6:       "IPv6",
			IPv4:       "IPv4",
			Success:    "正常",
			Error:      "错误",
			EndMessage: "秒后结束",
		}
	case "de":
		return Messages{
			CheckTitle: "Netzwerkprüfung",
			IPv6:       "IPv6",
			IPv4:       "IPv4",
			Success:    "OK",
			Error:      "Fehler",
			EndMessage: "Sekunden bis zum Beenden",
		}
	default:
		return Messages{
			CheckTitle: "Network Check",
			IPv6:       "IPv6",
			IPv4:       "IPv4",
			Success:    "OK",
			Error:      "Error",
			EndMessage: "seconds to exit",
		}
	}
}

func getLang() string {
	envVars := []string{"LC_ALL", "LC_MESSAGES", "LANG"}
	for _, env := range envVars {
		if value, ok := os.LookupEnv(env); ok && len(value) >= 2 {
			return strings.ToLower(value[:2])
		}
	}
	return "en"
}

func checkNetwork(url string) (bool, time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, 0
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	duration := time.Since(start)

	if err != nil {
		return false, duration
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, duration
	}

	return true, duration
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func ifColor(status bool, successMsg, errorMsg string) string {
	if isWindows() {
		if status {
			return successMsg
		}
		return errorMsg
	}

	if status {
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", successMsg)
	}
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", errorMsg)
}

func main() {
	msg := getMessages()

	fmt.Println("=========================")
	fmt.Println(msg.CheckTitle)
	fmt.Println("=========================")

	fmt.Printf("%s : 確認中···\r", msg.IPv6)
	ipv6Status, ipv6Duration := checkNetwork(ipv6Url)
	fmt.Printf("%s : %s (応答時間: %v)\n", msg.IPv6, ifColor(ipv6Status, msg.Success, msg.Error), ipv6Duration)

	fmt.Printf("%s : 確認中···\r", msg.IPv4)
	ipv4Status, ipv4Duration := checkNetwork(ipv4Url)
	fmt.Printf("%s : %s (応答時間: %v)\n", msg.IPv4, ifColor(ipv4Status, msg.Success, msg.Error), ipv4Duration)
	fmt.Println("=========================")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d%s\r", 10-i, msg.EndMessage)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}
