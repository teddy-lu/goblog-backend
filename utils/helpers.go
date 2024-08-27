package utils

import (
	mathrand "math/rand"
	"net"
	"time"
)

func GetUserIp() string {
	// Get the user's IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		// 排除回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func RandomString(length int) string {
	source := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[source.Intn(len(letters))]
	}
	return string(b)
}
