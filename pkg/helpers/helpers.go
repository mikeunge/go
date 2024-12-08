package helpers

import (
	"net"
	"regexp"
	"strconv"
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

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func ParseUint[T Uint](input string) T {
	var bitSize int

	switch any(T(0)).(type) {
	case uint8:
		bitSize = 8
	case uint16:
		bitSize = 16
	case uint32:
		bitSize = 32
	case uint64:
		bitSize = 64
	default: // Assume uint is platform-dependent (32 or 64 bits)
		bitSize = strconv.IntSize
	}

	parsedValue, err := strconv.ParseUint(input, 10, bitSize)
	if err != nil {
		return T(0)
	}
	return T(parsedValue)
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

func ParseInt[T Int](input string) T {
	var bitSize int

	switch any(T(0)).(type) {
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	default: // Assume uint is platform-dependent (32 or 64 bits)
		bitSize = strconv.IntSize
	}

	parsedValue, err := strconv.ParseInt(input, 10, bitSize)
	if err != nil {
		return T(0)
	}
	return T(parsedValue)
}
