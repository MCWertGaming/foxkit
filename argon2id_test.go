package foxkit

import "testing"

func TestHashString(t *testing.T) {
	testPass := "Test password 123456"

	// hashing password
	hash, err := CreateHash(&testPass)
	if err != nil {
		t.Errorf("Creating hash failed: %s", err.Error())
	}

	// decode hash
	params, err := DecodeHash(&hash)
	if err != nil {
		t.Errorf("Decoding hash failed: %s", err.Error())
	}

	// check the params list
	if params.Iterations != hashIterations {
		t.Errorf("Decoded hash iterations is %d instead of %d", params.Iterations, hashIterations)
	} else if params.KeyLength != hashKeyLength {
		t.Errorf("Decoded key length is %d instead of %d", params.KeyLength, hashKeyLength)
	} else if params.Memory != hashMemory {
		t.Errorf("Decoded hash memory is %d instead of %d", params.Memory, hashMemory)
	} else if params.Parallelism != hashParallelism {
		t.Errorf("Decoded key parallelism is %d instead of %d", params.Parallelism, hashParallelism)
	} else if params.SaltLength != hashSaltLength {
		t.Errorf("Decoded salt length is %d instead of %d", params.SaltLength, hashSaltLength)
	}

	// check the password
	match, err := ComparePasswordAndHash(&testPass, &hash)
	if err != nil {
		t.Errorf("Comparing the password and hash failed: %s", err.Error())
	} else if !match {
		t.Error("password and hashed password did not match")
	}
}

func TestHashedString(t *testing.T) {
	testPass := "Test password 1234"
	testHash := "$argon2id$v=19$m=65536,t=1,p=2$QfnqiXW2OjZhNE7aUxqynw$HrNvEP3PM6QDYyd+tjTrvH/oKSooNLIoq1Oj3PSqeyo"

	// Encode given hash
	params, err := DecodeHash(&testHash)
	if err != nil {
		t.Errorf("Decoding hash failed: %s", err.Error())
	} else if params.Parallelism != 2 {
		t.Errorf("Decoded hash parallelism was %d instead of %d", params.Parallelism, 2)
	} else if params.Memory != 65536 {
		t.Errorf("decoded hash Memory was %d instead of %d", params.Memory, 65536)
	} else if params.Iterations != 1 {
		t.Errorf("decoded hash iterations was %d instead of %d", params.Iterations, 1)
	} else if params.SaltLength != 16 {
		t.Errorf("decoded Salt length was %d instead of %d", params.SaltLength, 16)
	} else if params.KeyLength != 32 {
		t.Errorf("Decoded key length was %d instead of %d", params.KeyLength, 32)
	}

	match, err := ComparePasswordAndHash(&testPass, &testHash)
	if err != nil {
		t.Errorf("Comparing hash failed: %s", err.Error())
	} else if !match {
		t.Errorf("password and hash didn't match")
	}
}
