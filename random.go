package foxkit

import (
	"crypto/rand"
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
