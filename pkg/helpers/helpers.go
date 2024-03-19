package helpers

import (
	"net"
	"regexp"
)

func ValidateInputLength(input string, minLen, maxLen int) bool {
	if len(input) < minLen || len(input) > maxLen {
		return false
	}
	return true
}

func IsValidIp(host string) bool {
	if ip := net.ParseIP(host); ip == nil {
		return false
	}
	return true
}

func IsValidUrl(uri string) bool {
	if uri == "localhost" {
		return true
	}

	re := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	return re.MatchString(uri)
}
