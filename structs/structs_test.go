package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    user
		expected bool
	}{
		{name: "invalid_id", input: user{id: 0, age: 28, email: "den@yandex.ru", name: "Den"}, expected: false},
		{name: "invalid_age", input: user{id: 1, age: -1, email: "den@yandex.ru", name: "Den"}, expected: false},
		{name: "invalid_name", input: user{id: 1, age: 28, email: "den@yandex.ru", name: ""}, expected: false},
		{name: "valid_user", input: user{id: 1, age: 28, email: "den@yandex.ru", name: "Den"}, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.input.Validate()
			if actual != test.expected {
				t.Errorf("Expected %t, but got %t", test.expected, actual)
			}
		})
	}

}

func TestTotal(t *testing.T) {
	p1 := product{id: 1, name: "1", price: 15}
	p2 := product{id: 2, name: "2", price: 35}
	o := order{id: 1, price: 10, products: []product{p1, p2}}
	if o.Total() != 50 {
		t.Error("Expected ", 50)
	}
}
