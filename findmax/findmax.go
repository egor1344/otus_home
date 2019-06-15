package findmax

import (
	"reflect"
)

// Max - func return index max value
func Max(slice interface{}, less func(i, j int) bool) int {
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	maxValueIndex := findMax(less, length)
	return maxValueIndex
}

func findMax(lessFunc func(i, j int) bool, len int) int {
	maxValueIndex := 0
	for i := 0; i < (len - 1); i++ {
		if !lessFunc(maxValueIndex, i+1) {
			maxValueIndex = i + 1
		}
	}
	return maxValueIndex
}
