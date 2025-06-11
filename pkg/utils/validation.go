package utils

import (
	"regexp"
)

// IsValidEmail checks if the input string is a valid email format.
func IsValidEmail(email string) bool {
	// This regex is RFC 5322 compliant enough for general use.
	const emailRegexPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}
