package pipeline

import (
	"context"
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
