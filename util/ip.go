package util

import (
	"errors"
	"io"
	"net"
	"net/http"
)

// GetLANIP 获取局域网IP
func GetLANIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsPrivate() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can not find the LAN ip address!")
}

// GetWANIP 获取公网IP
func GetWANIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	return string(content), nil
}
