package pipeline

import (
	"context"
)

func Generate(
	ctx context.Context,
	numbers []int,
) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for _, num := range numbers {
			select {
			case channel <- num:
			case <-ctx.Done():
				return
			}
		}
	}()

	return channel
}

func Square(
	ctx context.Context,
	input <-chan int,
) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for {
			select {
			case num, ok := <-input:
				if !ok {
					return
				}
				select {
				case channel <- num * num:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return channel
}

func FilterEven(
	ctx context.Context,
	input <-chan int,
) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for {
			select {
			case num, ok := <-input:
				if !ok {
					return
				}
				if num%2 != 0 {
					continue
				}
				select {
				case channel <- num:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return channel
}

func Collect(
	ctx context.Context,
	input <-chan int,
) ([]int, error) {
	result := make([]int, 0, 5)

	for {
		select {
		case num, ok := <-input:
			if !ok {
				return result, nil
			}
			result = append(result, num)
		case <-ctx.Done():
			return []int{}, ctx.Err()
		}
	}
}
