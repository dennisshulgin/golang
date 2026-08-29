package concurrencycontext

import (
	"context"
	"testing"
	"time"
)

func TestWaitForValue(t *testing.T) {
	ctx := context.Background()
	ch := make(chan int, 1)
	ch <- 42
	value, err := WaitForValue(ctx, ch)

	if value != 42 || err != nil {
		t.Error("Unexpected error")
	}
}

func TestWaitForValueTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	ch := make(chan int)
	value, err := WaitForValue(ctx, ch)

	if value != 0 || err != context.DeadlineExceeded {
		t.Error("Unexpected value")
	}
}

func TestWaitForValueCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	cancel()
	value, err := WaitForValue(ctx, ch)

	if value != 0 || err != context.Canceled {
		t.Error("Unexpected value")
	}
}

func TestProcessValues(t *testing.T) {
	ctx := context.Background()
	input := make(chan int, 1)
	output := make(chan int, 1)
	input <- 5
	close(input)
	err := ProcessValues(ctx, input, output)

	if err != nil {
		t.Error("Unexpected value")
	}

	num := <-output
	if num != 25 {
		t.Error("Expected 25")
	}
}

func TestProcessValuesTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	input := make(chan int, 1)
	output := make(chan int, 1)
	err := ProcessValues(ctx, input, output)

	if err != context.DeadlineExceeded {
		t.Error("Unexpected value")
	}
}
