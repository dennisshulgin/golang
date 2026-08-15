package channels

import (
	"testing"
	"sort"
	"slices"
)

func TestGenerateSquares(t *testing.T) {
	tests := []struct {
		name string
		input []int
		sortedExpected []int
	} {
		{name: "empty", input: []int{}, sortedExpected: []int{}},
		{name: "calculated", input: []int{1, 2, 3, 4, 5}, sortedExpected: []int{1, 4, 9, 16, 25}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GenerateSquares(test.input)
			sort.Ints(actual)

			if !slices.Equal(actual, test.sortedExpected) {
				t.Errorf("Expected %v, but got %v", test.sortedExpected, actual)
			}
		})
	}
}