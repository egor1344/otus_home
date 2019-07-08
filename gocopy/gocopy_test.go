package gocopy


import (
	"testing"
	)

func TestCopyFile(t *testing.T) {
	GoCopy("/home/mrk/Pictures/1.MP4", "/home/mrk/Pictures/1_copy.MP4", 1024, 1024)
}
