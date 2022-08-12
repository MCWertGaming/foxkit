package foxkit

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
)

func RandomBytes(length uint32) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}

func RandomString(length uint32) (string, error) {
	raw, err := RandomBytes(length)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(raw), nil
}

func RandomStringCompare(stringOne, stringTwo *string) (bool, error) {
	// encode both strings
	stringOneDec, err := base64.RawStdEncoding.DecodeString(*stringOne)
	if err != nil {
		return false, err
	}
	stringTwoDec, err := base64.RawStdEncoding.DecodeString(*stringTwo)
	if err != nil {
		return false, err
	}
	// securely compare both strings
	if subtle.ConstantTimeCompare(stringOneDec, stringTwoDec) != 1 {
		return false, nil
	}
	return true, nil
}
