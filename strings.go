package foxkit

import (
	"net/mail"
	"strings"
)

// returns false if a given parameter is false, minSize, maxSize, ascii
// TODO: Test
func CheckStringFull(testString string, minLength, maxLength uint32) (minSize bool, maxSize bool, ascii bool) {
	if strings.Count(testString, "") > int(maxLength+1) {
		minSize = false
	} else {
		maxSize = true
	}
	if strings.Count(testString, "") < int(minLength+1) {
		maxSize = false
	} else {
		maxSize = true
	}
	return minSize, maxSize, IsASCII(testString)
}

// returns true if all parameters are true
func CheckString(testString string, minLength, maxLength uint32, asciiOnly bool) bool {
	minSize, maxSize, ascii := CheckStringFull(testString, minLength, maxLength)
	if minSize && maxSize && ((asciiOnly && ascii) || !asciiOnly) {
		return true
	} else {
		return false
	}
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
