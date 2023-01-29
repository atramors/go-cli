package helpers

import (
	"strings"
)

// Checking if provided string a number.
func ArgIsNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	digits := "1234567890"
	for _, char := range s {
		if !strings.ContainsAny(string(char), digits) {
			return false
		}
	}
	return true
}
