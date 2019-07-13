package gocopy

import (
	"testing"
)

func TestCopyFile(t *testing.T) {
	err := GoCopy("/home/mrk/Videos/1.txt", "/home/mrk/Videos/1_copy.txt", 2048, 1024)
	if err != nil {
		t.Error("Wrong result ", err)
	}
	err = GoCopy("/home/mrk/Videos/1.txt", "/home/mrk/Videos/1_copy.txt", 2048, 3000)
	if err == nil {
		t.Error("Wrong result ")
	}
}
