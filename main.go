package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	ipv6Url = "https://check-ipv6.sangmin.eu.org/check.json"
	ipv4Url = "https://check-ipv4.sangmin.eu.org/check.json"
	timeout = 5 * time.Second
)

func checkNetwork(url string) (bool, time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, 0
	}

	start := time.Now() // 요청 시작 시간 기록
	resp, err := http.DefaultClient.Do(req)
	duration := time.Since(start) // 요청 종료 후 시간 측정

	if err != nil {
		return false, duration
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, duration
	}

	return true, duration
}

func main() {
	fmt.Println("=========================")
	fmt.Println("ネットチェック")
	fmt.Println("=========================")

	fmt.Printf("IPv6 : 確認中···\r")
	ipv6Status, ipv6Duration := checkNetwork(ipv6Url)
	fmt.Printf("IPv6 : %s (応答時間: %v)\n", ifColor(ipv6Status, "正常", "エラー"), ipv6Duration)

	fmt.Printf("IPv4 : 確認中···\r")
	ipv4Status, ipv4Duration := checkNetwork(ipv4Url)
	fmt.Printf("IPv4 : %s (応答時間: %v)\n", ifColor(ipv4Status, "正常", "エラー"), ipv4Duration)
	fmt.Println("=========================")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d秒後に終了\r", 10-i)
    	time.Sleep(1 * time.Second)
	}
}

func ifColor(status bool, successMsg, errorMsg string) string {
	if status {
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", successMsg)
	}
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", errorMsg)
}
