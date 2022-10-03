package foxkit

import (
	"testing"
)

func TestCheckString(t *testing.T) {
	// Test if the function notices too small strings
	if CheckString("test", 0, 0, true) {
		t.Error("size 0,0 test failed")
	} else if CheckString("test", 0, 1, true) { // string to small
		t.Error("size 0,1 test failed")
	} else if CheckString("test", 0, 2, true) { // string to small
		t.Error("size 0,2 test failed")
	} else if CheckString("test", 0, 3, true) { // string to small
		t.Error("size 0,3 test failed")
	} else if !CheckString("test", 0, 4, true) { // the string is the right size
		t.Error("size 0,4 test failed")
	} else if !CheckString("test", 0, 5, true) { // the string is the right size
		t.Error("size 0,5 test failed")
	} else if !CheckString("test", 4, 4, true) { // the right size
		t.Error("size 4,4 test failed")
	} else if CheckString("test", 5, 10, true) { // to small
		t.Error("size 0,0 test failed")
	} else if CheckString("Test oLoLdsawer", 8, 20, true) { // right size
		t.Error("size 8,20 test failed")
	} else if CheckString("Test lololol1234", 8, 20, true) { // right size, no unicode
		t.Error("no unicode test failed")
	} else if !CheckString("A unicode char: ♔", 3, 20, false) { // check if unicode is accepted
		t.Error("unicode allowed test failed")
	} else if CheckString("A unicode char: ♔", 3, 20, true) { // detect unicode
		t.Error("unicode was not detected")
	} else if !CheckString("Die gestielten Blüten der Apfelbäume stehen einzeln oder in doldigen schirmrispigen Blütenständen.", 10, 100, false) {
		t.Error("unicode was not allowed in german sentence")
	} else if CheckString("Die gestielten Blüten der Apfelbäume stehen einzeln oder in doldigen schirmrispigen Blütenständen.", 10, 100, true) {
		t.Error("unicode was not detected in german sentence")
	}
}

func TestCheckEmail(t *testing.T) {
	if !CheckEmail("test@example.invalid") {
		t.Error("test@example.invalid not accepted")
	} else if CheckEmail("") {
		t.Error("Empty email field was allowed")
	}
}
