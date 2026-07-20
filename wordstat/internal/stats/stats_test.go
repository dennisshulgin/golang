package stats

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestGetWords(t *testing.T) {
	s := "Denis denis! Hello my friend! I live in New-York... Ready to  work. "
	actualWords := GetWords(s)
	expectedWords := []string{"denis", "denis", "hello", "my", "friend", "i", "live", "in", "new", "york", "ready", "to", "work"}
	fmt.Println(actualWords)
	if !reflect.DeepEqual(actualWords, expectedWords) {
		t.Error("Slices are not equal")
	}
}

func TestReadAsString(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected string
	}{
		{name: "string_with_n", input: bytes.NewBufferString("Denis denis \n qwerty"), expected: "Denis denis "},
		{name: "string_without_n", input: bytes.NewBufferString("Denis denis qwerty"), expected: "Denis denis qwerty"},
		{name: "empty_string", input: bytes.NewBufferString(""), expected: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := ReadAsString(test.input)
			if result != test.expected {
				t.Errorf("Got %s, but expected %s", result, test.expected)
			}
		})
	}
}
