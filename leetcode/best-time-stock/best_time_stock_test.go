package besttimestock

import "testing"

func TestBestTimeStock(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "default", input: []int{7, 5, 1, 8, 10}, expected: 9},
		{name: "empty", input: []int{}, expected: 0},
		{name: "one_element", input: []int{1}, expected: 0},
		{name: "decreased_sequence", input: []int{7, 6, 5, 4}, expected: 0},
		{name: "increased_sequence", input: []int{1, 2, 3, 4}, expected: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := MaxProfit(test.input)

			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
