package main

import (
	"testing"
)

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "empty_slice", input: []int{}, expected: []int{}},
		{name: "without_duplicates", input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
		{name: "some_elements", input: []int{1, 2, 3, 1, 2}, expected: []int{1, 2, 3}},
		{name: "same_elements", input: []int{1, 1, 1, 1, 1, 1}, expected: []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Deduplicate(test.input)
			if !compareSlices(output, test.expected) {
				t.Errorf("Expected %d, but got %d", test.expected, output)
			}
		})
	}

}

func compareSlices(a, b []int) bool {
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
