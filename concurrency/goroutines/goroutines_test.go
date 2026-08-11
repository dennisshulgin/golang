package goroutines

import (
	"slices"
	"sort"
	"testing"
)

func TestRunWorkers(t *testing.T) {
	actual := RunWorkers(5)
	sort.Ints(actual)
	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}