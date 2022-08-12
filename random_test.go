package foxkit

import "testing"

func TestRandomString(t *testing.T) {
	random, err := RandomString(20)
	if err != nil {
		t.Errorf("Generating random string failed: %s", err.Error())
	}

	// check if a single string matches
	match, err := RandomStringCompare(&random, &random)
	if err != nil {
		t.Errorf("Generating random string failed: %s", err.Error())
	}
	if !match {
		t.Error("The same generated string didn't match")
	}

	// check if two separate strings match
	randomOne, err := RandomString(30)
	if err != nil {
		t.Errorf("Generating random string failed: %s", err.Error())
	}
	randomTwo, err := RandomString(30)
	if err != nil {
		t.Errorf("Generating random string failed: %s", err.Error())
	}
	match, err = RandomStringCompare(&randomOne, &randomTwo)
	if err != nil {
		t.Errorf("Generating random string failed: %s", err.Error())
	}
	if match {
		t.Errorf("Two generated strings matched: %s and %s", randomOne, randomTwo)
	}
}
