package doublelinkedlist

import (
	"testing"
		// "log"
	)

func TestDuobleLinkedList(t *testing.T) {
	testCase := []int{0,1,2}
	maxFunc := func (i, j int) bool {
		return testCase[i] > testCase[j]
	}
	if resultIndex := Max(testCase, maxFunc); (resultIndex != 2) {
		t.Error("Wrong result ", resultIndex)
	}
}