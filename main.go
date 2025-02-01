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

func checkNetwork(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	fmt.Println("=========================")
	fmt.Println("ネットチェック")
	fmt.Println("=========================")

	ipv6Status := checkNetwork(ipv6Url)
	ipv4Status := checkNetwork(ipv4Url)

	fmt.Printf("IPv6 : %s\n", ifColor(ipv6Status, "正常", "エラー"))
	fmt.Printf("IPv4 : %s\n", ifColor(ipv4Status, "正常", "エラー"))
	fmt.Println("=========================")

	time.Sleep(10 * time.Second)
}

func ifColor(status bool, successMsg, errorMsg string) string {
	if status {
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", successMsg)
	}
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", errorMsg)
}
