package main


import "testing"

func TestHelloWorld(t *testing.T) {
	actual := GetQuote()
	expected := "Don't communicate by sharing memory, share memory by communicating."
	if actual != expected {
		t.Errorf("got %q, but expected %q", actual, expected)
	}
}
