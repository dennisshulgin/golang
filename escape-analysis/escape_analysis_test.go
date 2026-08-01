package escapeanalysis

import (
	"testing"
)

var userPointer *User
var userValue User

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum(1, 2)
	}
}

func BenchmarkNewUserValue(b *testing.B) {
	for b.Loop() {
		userValue = NewUserValue()
	}
}

func BenchmarkNewUserPointer(b *testing.B) {
	for b.Loop() {
		userPointer = NewUserPointer()
	}
}