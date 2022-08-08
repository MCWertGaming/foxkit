package foxkit

import (
	"strings"
	"testing"
)

// test if UUID generation works
func TestUUID(t *testing.T) {
	uuid := getUUID()
	if strings.Count(uuid, "") != 37 {
		t.Errorf("UUID not 37 characters long, instead %d", strings.Count(uuid, ""))
	}
}

func TestUUIDPool(t *testing.T) {
	setupUUID()
	uuid := getUUID()

	if strings.Count(uuid, "") != 37 {
		t.Errorf("UUID not 37 characters long, instead %d", strings.Count(uuid, ""))
	}
}
