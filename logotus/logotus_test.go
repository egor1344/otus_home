package logotus

import (
	"testing"
	"os"
	)

func TestDuobleLinkedItem(t *testing.T) {
	testHwAccepted := HwAccepted{1,2}
	var i OtusEvent
	i = &testHwAccepted
	LogOtusEvent(i, os.Stdout)
	testHwSubmitted := HwSubmitted{1,"code", "Comment"}
	i = &testHwSubmitted
	LogOtusEvent(i, os.Stdout)	
}
