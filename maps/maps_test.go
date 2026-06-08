package main

import "testing"
import "maps"

func TestCountWords(t *testing.T) {
	actual := CountWords("1 1 1 2 3 4")
	expected := map[string]int{
		"1": 3,
		"2": 1,
		"3": 1,
		"4": 1,
	}
	areEqual := maps.Equal(actual, expected)
	if !areEqual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
