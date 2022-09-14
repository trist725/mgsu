package util

import "testing"

func TestGetLANIP(t *testing.T) {
	t.Log(GetLANIP())
}

func TestGetWANIP(t *testing.T) {
	ip, _ := GetWANIP()
	t.Log(ip)
}
