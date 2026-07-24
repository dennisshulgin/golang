package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "1", input: "A man, a plan, a canal: Panama", expected: true},
		{name: "2", input: "race a car", expected: false},
		{name: "3", input: " ", expected: true},
		{name: "4", input: "0P", expected: false},
		{name: "5", input: "А роза упала на лапу Азора", expected: true},
		{name: "6", input: "No lemon, no melon", expected: true},
		{name: "7", input: "", expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsPalindrome(test.input)

			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
