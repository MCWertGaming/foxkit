package foxkit

import (
	"net/mail"
	"strings"
)

// returns true if the string matches the given schema
func CheckString(testString string, minLength, maxLength uint8, asciiOnly bool) bool {
	if strings.Count(testString, "") > int(maxLength+1) || strings.Count(testString, "") < int(minLength+1) {
		return false
	} else if asciiOnly && !IsASCII((testString)) {
		return false
	}
	return true
}

// returns true if the string only contains ascii characters
func IsASCII(s string) bool {
	for _, c := range s {
		if c > 126 || c < 33 { // ascii goes from 33 to 126
			return false
		}
	}
	return true
}

// returns true if the given string is an email
func CheckEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}
