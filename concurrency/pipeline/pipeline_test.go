package pipeline

import (
	"context"
	"errors"
	"slices"
	"testing"
)

func TestPipeline(t *testing.T) {
	ctx := context.Background()
	expected := []int{4, 16}

	generated := Generate(ctx, []int{1, 2, 3, 4})
	squared := Square(ctx, generated)
	filtered := FilterEven(ctx, squared)
	result, _ := Collect(ctx, filtered)

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestEmpty(t *testing.T) {
	ctx := context.Background()
	expected := []int{}

	generated := Generate(ctx, []int{})
	squared := Square(ctx, generated)
	filtered := FilterEven(ctx, squared)
	result, err := Collect(ctx, filtered)

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	if err != nil {
		t.Error("Expected nil, but got not nil")
	}
}

func TestManyNumbers(t *testing.T) {
	ctx := context.Background()
	numsCount := 1000
	nums := make([]int, 0, numsCount)

	for i := 0; i < numsCount; i++ {
		nums = append(nums, i)
	}

	expected := make([]int, 0, 1000)

	for i := 0; i < numsCount; i++ {
		if (i*i)%2 == 0 {
			expected = append(expected, i*i)
		}
	}

	generated := Generate(ctx, nums)
	squared := Square(ctx, generated)
	filtered := FilterEven(ctx, squared)
	result, err := Collect(ctx, filtered)

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	if err != nil {
		t.Error("Expected nil, but got not nil")
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	generated := Generate(ctx, []int{1, 2, 3, 4})
	squared := Square(ctx, generated)
	filtered := FilterEven(ctx, squared)
	_, err := Collect(ctx, filtered)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected %v, but got %v", context.Canceled, err)
	}
}
