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

func TestUUIDCollision(t *testing.T) {
	if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	}
}

func TestUUIDPoolCollision(t *testing.T) {
	SetupUUID()
	if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	} else if GetUUID() == GetUUID() {
		t.Error("Two UUIDs collided, this should be impossible")
	}
}
