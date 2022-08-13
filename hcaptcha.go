package foxkit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Credit      bool      `json:"credit"`
	ErrorCode   []string  `json:"error-codes"`
}

func HCaptchaValid(secret, response *string, hostname string, validFor time.Duration) (bool, error) {
	// get response from hcaptcha
	resp, err := http.Post("https://hcaptcha.com/siteverify", "application/x-www-form-urlencoded; charset=utf-8", strings.NewReader(fmt.Sprintf("secret=%s&response=%s", *secret, *response)))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// parse response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// put response into struct
	var captcha HCaptchaResponse
	err = json.Unmarshal([]byte(bodyBytes), &captcha)
	if err != nil {
		return false, err
	}

	// check response
	if !captcha.Success {
		if captcha.ErrorCode != nil {
			return false, fmt.Errorf("captcha is invalid, error code: %v", captcha.ErrorCode)
		} else {
			return false, fmt.Errorf("captcha is invalid without error code")
		}
	} else if captcha.Hostname != hostname {
		return false, fmt.Errorf("wrong hostname for captcha challenge, got %s instead of %s", captcha.Hostname, hostname)
	} else if !captcha.ChallengeTS.After(time.Now().Add(-validFor)) {
		return false, fmt.Errorf("challenge to old therefore invalid")
	}

	// challenge valid
	return true, nil
}
