package foxkit

import (
	"strings"
	"testing"
)

// test if UUID generation works
func TestUUID(t *testing.T) {
	uuid := GetUUID()
	if strings.Count(uuid, "") != 33 {
		t.Errorf("UUID not 37 characters long, instead %d", strings.Count(uuid, ""))
	}
}

func TestUUIDPool(t *testing.T) {
	SetupUUID()
	uuid := GetUUID()

	if strings.Count(uuid, "") != 33 {
		t.Errorf("UUID not 37 characters long, instead %d", strings.Count(uuid, ""))
	}
}
