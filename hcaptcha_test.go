package foxkit

import (
	"testing"
	"time"
)

func TestHCaptcha(t *testing.T) {
	secretKey := "0x0000000000000000000000000000000000000000"
	response := "10000000-aaaa-bbbb-cccc-000000000001"

	var captcha HCaptchaResponse
	err := captcha.Get(&secretKey, &response)
	if err != nil {
		t.Errorf("Retrieving captcha response failed: %s", err.Error())
	}
	valid, err := captcha.Valid("dummy-key-pass", time.Hour)
	if err != nil {
		t.Errorf("Failed to validate HCaptcha: %s", err.Error())
	} else if !valid {
		t.Errorf("Sample captcha not valid")
	}
}
