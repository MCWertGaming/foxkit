package foxkit

import (
	"strings"

	"github.com/google/uuid"
)

// enables using the random pool, might be insecure and isn't thread safe
func SetupUUID() {
	uuid.EnableRandPool()
}

// returns a randomly generated UUID (v4) without hyphen
func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
