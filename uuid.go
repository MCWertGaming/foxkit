package foxkit

import "github.com/google/uuid"

// enables using the random pool, might be insecure and isn't thread safe
func SetupUUID() {
	uuid.EnableRandPool()
}

// returns a randomly generated UUID (v4)
func GetUUID() string {
	return uuid.New().String()
}
