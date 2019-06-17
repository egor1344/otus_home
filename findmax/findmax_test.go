package findmax

import (
	"testing"
		// "log"
	)

func TestFinMax(t *testing.T) {
	testCase := []int{0,1,2}
	testInterface := make([]interface{}, len(testCase))
	for i, j:= range testCase {
		testInterface[i] = j
	}
	maxFunc := func (i, j int) bool {
		return testCase[i] > testCase[j]
	}
	if resultIndex, err := Max(testInterface, maxFunc); (resultIndex != 2 && err==nil) {
		t.Error("Wrong result ", resultIndex)
	}
	testCase = []int{3,1,2}
	testInterface = make([]interface{}, len(testCase))
	for i, j:= range testCase {
		testInterface[i] = j
	}
	if resultIndex, err := Max(testInterface, maxFunc); (resultIndex != 0 && err==nil) {
		t.Error("Wrong result ", resultIndex)
	}
	testCase = []int{3,4,2}
	testInterface = make([]interface{}, len(testCase))
	for i, j:= range testCase {
		testInterface[i] = j
	}
	if resultIndex, err := Max(testInterface, maxFunc); (resultIndex != 1 && err==nil) {
		t.Error("Wrong result ", resultIndex)
	}
	testCase = []int{}
	testInterface = make([]interface{}, len(testCase))
	for i, j:= range testCase {
		testInterface[i] = j
	}
	if resultIndex, err := Max(testInterface, maxFunc); (resultIndex != 0 && err==nil) {
		t.Error("Wrong result ", resultIndex)
	}
	testStringCase := []string{"1", "22", "333"}
	testInterface = make([]interface{}, len(testCase))
	for i, j:= range testCase {
		testInterface[i] = j
	}
	maxStringFunc := func (i, j int) bool {
		return testStringCase[i] > testStringCase[j]
	}
	if resultIndex, err := Max(testInterface, maxStringFunc); (resultIndex != 2 && err==nil) {
		t.Error("Wrong result ", resultIndex)
	}
}