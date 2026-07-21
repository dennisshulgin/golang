package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "default", nums: []int{1, 2, 3, 4}, target: 5, expected: []int{1, 2}},
		{name: "all_equals", nums: []int{1, 1, 1, 1}, target: 2, expected: []int{0, 1}},
		{name: "not_found", nums: []int{1, 2, 3, 4}, target: 10, expected: []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := TwoSum(test.nums, test.target)

			if !areEqualSlices(actual, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}

func areEqualSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
