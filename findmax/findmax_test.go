package findmax

import (
	"testing"
		// "log"
	)

func TestFinMax(t *testing.T) {
	testCase := []int{0,1,2}
	maxFunc := func (i, j int) bool {
		return testCase[i] > testCase[j]
	}
	if resultIndex := Max(testCase, maxFunc); (resultIndex != 2) {
		t.Error("Wrong result ", resultIndex)
	}
	testCase = []int{3,1,2}
	if resultIndex := Max(testCase, maxFunc); (resultIndex != 0) {
		t.Error("Wrong result ", resultIndex)
	}
	testCase = []int{3,4,2}
	if resultIndex := Max(testCase, maxFunc); (resultIndex != 1) {
		t.Error("Wrong result ", resultIndex)
	}
	testStringCase := []string{"1", "22", "333"}
	maxStringFunc := func (i, j int) bool {
		return testStringCase[i] > testStringCase[j]
	}
	if resultIndex := Max(testStringCase, maxStringFunc); (resultIndex != 2) {
		t.Error("Wrong result ", resultIndex)
	}
}