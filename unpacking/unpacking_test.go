package unpacking

import "testing"

func TestUnpacking(t *testing.T) {
	if result := Unpacking("a4bc2d5e"); result != "aaaabccddddde" {
		t.Error("Wrong result ", result, " was expected aaaabccddddde")
	}
	if result := Unpacking("abcd"); result != "abcd" {
		t.Error("Wrong result ", result, " was expected abcd")
	}
	if result := Unpacking("45"); result != "" {
		t.Error("Wrong result ", result, " was expected blank string")
	}
	if result := Unpacking("qwe\\4\\5"); result != "qwe45" {
		t.Error("Wrong result ", result, " was expected qwe45")
	}
	if result := Unpacking("qwe\\45"); result != "qwe44444" {
		t.Error("Wrong result ", result, " was expected qwe44444")
	}
	if result := Unpacking("qwe\\\\5"); result != "qwe\\\\\\\\\\" {
		t.Error("Wrong result ", result, " was expected qwe\\\\\\\\\\")
	}
	
}