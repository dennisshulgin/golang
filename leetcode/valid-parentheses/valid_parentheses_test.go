package validparentheses

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "empty_stirng", input: "", expected: true},
		{name: "one_parenthese", input: "(", expected: false},
		{name: "default", input: "({})", expected: true},
		{name: "invalid_positions", input: "{(})", expected: false},
		{name: "invalid_count", input: "((()", expected: false},
		{name: "different_parentheses", input: "{)", expected: false},
		{name: "correct", input: "[]{}()((())){}{}", expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsValid(test.input)

			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}
