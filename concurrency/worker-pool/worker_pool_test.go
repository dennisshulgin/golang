package workerpool

import (
	"testing"
	"sort"
	"slices"
)

func TestProcessJobs(t *testing.T) {
	actual := ProcessJobs([]int{1, 2, 3, 4, 5}, 1)
	sort.Ints(actual)
	expected := []int{1, 4, 9, 16, 25}
	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestProcessJobsEmpty(t *testing.T) {
	actual := ProcessJobs([]int{}, 1)
	sort.Ints(actual)
	expected := []int{}
	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
