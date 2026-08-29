package concurrencycontext

import (
	"context"
)

func WaitForValue(
	ctx context.Context,
	ch <-chan int,
) (int, error) {
	select {
	case msg := <-ch:
		return msg, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func ProcessValues(
	ctx context.Context,
	input <-chan int,
	output chan<- int,
) error {
	for {
		select {
		case num, ok := <-input:
			if !ok {
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- num * num:
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
