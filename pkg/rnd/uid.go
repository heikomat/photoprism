package rnd

import (
	"strconv"
	"time"
)

// PPID returns a unique id with prefix as string.
func PPID(prefix byte) string {
	result := make([]byte, 0, 16)
	result = append(result, prefix)
	result = append(result, strconv.FormatInt(time.Now().UTC().Unix(), 36)[0:6]...)
	result = append(result, Token(9)...)

	return string(result)
}

// IsPPID returns true if string is a unique id as generated by PhotoPrism.
func IsPPID(s string, prefix byte) bool {
	if len(s) != 16 {
		return false
	}

	return s[0] == prefix && IsLowerAlnum(s)
}

// IsHex returns true if the string only contains hex numbers, dashes and letters without whitespace.
func IsHex(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 48 || r > 57) && (r < 97 || r > 102) && (r < 65 || r > 90) && r != 45{
			return false
		}
	}

	return true
}

// IsLowerAlnum returns true if the string only contains alphanumeric ascii chars without whitespace.
func IsLowerAlnum(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 48 || r > 57) && (r < 97 || r > 122) {
			return false
		}
	}

	return true
}

// IsUID returns true if string is a seemingly unique id.
func IsUID(s string, prefix byte) bool {
	// Regular UUID.
	if len(s) == 36 && IsHex(s) {
		return true
	}

	// Not a known UID format.
	if len(s) != 16 {
		return false
	}

	return IsPPID(s, prefix)
}
