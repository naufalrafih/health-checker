package main

import (
	"fmt"
	"net"
	"time"
)

func checkStatus(domain string, port string) string {
	url := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", url, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable\nError: %v", url, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable\nFrom: %v\nTo: %v", url, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
