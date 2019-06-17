package findmax

import (
	"errors"
)

// Max - func return index max value
func Max(slice []interface{}, less func(i, j int) bool) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice len = 0")
	}
	var maxValueIndex int
	for i := range slice {
		if !less(maxValueIndex, i) {
			maxValueIndex = i
		}
	}
	return maxValueIndex, nil
}
