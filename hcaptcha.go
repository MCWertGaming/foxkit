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

// retrieves the HCaptcha response
func (hresp *HCaptchaResponse) Get(secret, response *string) error {
	// get response from hcaptcha
	resp, err := http.Post("https://hcaptcha.com/siteverify", "application/x-www-form-urlencoded; charset=utf-8", strings.NewReader(fmt.Sprintf("secret=%s&response=%s", *secret, *response)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// parse response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// put response into struct
	err = json.Unmarshal([]byte(bodyBytes), &hresp)
	if err != nil {
		return err
	}

	// everything worked
	return nil
}

// checks if the hcaptcha response is valid, returns error if the response is invalid
func (captcha *HCaptchaResponse) Valid(hostname string, validFor time.Duration) (bool, error) {

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
