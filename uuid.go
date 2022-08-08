package foxkit

import "github.com/google/uuid"

// enables using the random pool, might be insecure and isn't thread safe
func setupUUID() {
	uuid.EnableRandPool()
}

// returns a randomly generated UUID (v4)
func getUUID() string {
	return uuid.New().String()
}
